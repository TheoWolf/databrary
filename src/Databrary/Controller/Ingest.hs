{-# LANGUAGE OverloadedStrings #-}
module Databrary.Controller.Ingest
  ( viewIngest
  , postIngest
  , detectParticipantCSV
  , runParticipantUpload
  ) where

import Control.Arrow (right)
import Control.Monad (unless)
import Control.Monad.IO.Class (liftIO)
import qualified Data.Attoparsec.ByteString as ATTO
import qualified Data.ByteString as BS
import qualified Data.ByteString.Char8 as BSC
import qualified Data.ByteString.Lazy as BSL
import qualified Data.Csv as CSV
import qualified Data.Csv.Parser as CSVP
import qualified Data.HashMap.Strict as HMP
import qualified Data.List as L
import Data.Maybe (catMaybes)
import qualified Data.Map as MAP
import Data.Map (Map)
import Data.Monoid ((<>))
import Data.Text (Text)
import qualified Data.Text as T
import qualified Data.Text.Encoding as TE
import qualified Data.Text.Lazy.Encoding as TLE
import qualified Data.Text.Lazy as TL
import qualified Data.Vector as V
import Data.Word (Word64)
import Network.HTTP.Types (badRequest400)
import Network.Wai.Parse (FileInfo(..))
import System.Posix.FilePath (takeExtension)

import qualified Databrary.JSON as JSON
import Databrary.Ops
import Databrary.Has
import Databrary.Model.Id
import Databrary.Model.Permission
import Databrary.Model.Volume
import Databrary.Model.Container
import Databrary.Model.Metric (ParticipantFieldMapping(..), Metric(..))
import Databrary.Model.VolumeMetric (lookupParticipantFieldMapping)
import Databrary.Model.Record
import Databrary.Model.Ingest -- (requiredColumnsPresent, headerMappingJSON, HeaderMappingEntry(..))
import Databrary.Ingest.Action
import Databrary.Ingest.JSON
import Databrary.HTTP.Path.Parser
import Databrary.HTTP.Form.Deform
import Databrary.Action.Route
import Databrary.Action
import Databrary.Controller.Paths
import Databrary.Controller.Permission
import Databrary.Controller.Form
import Databrary.Controller.Volume
import Databrary.View.Ingest
import Databrary.View.Form (FormHtml)

viewIngest :: ActionRoute (Id Volume)
viewIngest = action GET (pathId </< "ingest") $ \vi -> withAuth $ do
  checkMemberADMIN
  s <- focusIO getIngestStatus
  v <- getVolume PermissionEDIT vi
  peeks $ blankForm . htmlIngestForm v s

postIngest :: ActionRoute (Id Volume)
postIngest = multipartAction $ action POST (pathId </< "ingest") $ \vi -> withAuth $ do
  checkMemberADMIN
  s <- focusIO getIngestStatus
  v <- getVolume PermissionEDIT vi
  a <- runFormFiles [("json", 16*1024*1024)] (Just $ htmlIngestForm v s) $ do
    csrfForm
    abort <- "abort" .:> deform
    abort ?!$> (,,)
      <$> ("run" .:> deform)
      <*> ("overwrite" .:> deform)
      <*> ("json" .:> do
              (fileInfo :: FileInfo JSON.Value) <- deform
              (deformCheck
                   "Must be JSON."
                   (\f -> fileContentType f `elem` ["text/json", "application/json"] || takeExtension (fileName f) == ".json")
                   fileInfo))
  r <- maybe
    (True <$ focusIO abortIngest)
    (\(r,o,j) -> runIngest $ right (map (unId . containerId . containerRow)) <$> ingestJSON v (fileContent j) r o)
    a
  unless r $ result $ response badRequest400 [] ("failed" :: String)
  peeks $ otherRouteResponse [] viewIngest (volumeId $ volumeRow v)

maxCsvSize :: Word64
maxCsvSize = 16*1024*1024

-- TODO: maybe put csv file save/retrieve in Databrary.Store module
detectParticipantCSV :: ActionRoute (Id Volume)
detectParticipantCSV = action POST (pathJSON >/> pathId </< "detectParticipantCSV") $ \vi -> withAuth $ do
    v <- getVolume PermissionEDIT vi
    reqCtxt <- peek
    csvFileInfo <-
      -- TODO: is Nothing okay here?
      runFormFiles [("file", maxCsvSize)] (Nothing :: Maybe (RequestContext -> FormHtml TL.Text)) $ do
          csrfForm
          fileInfo :: (FileInfo TL.Text) <- "file" .:> deform
          return fileInfo
    let uploadFileContents = (BSL.toStrict . TLE.encodeUtf8 . fileContent) csvFileInfo
        uploadFileName = (BSC.unpack . fileName) csvFileInfo  -- TODO: add prefix to filename
    case ATTO.parseOnly (CSVP.csvWithHeader CSVP.defaultDecodeOptions) uploadFileContents of
        Left err -> do
            liftIO (print ("csv parse error", err))
            pure (forbiddenResponse reqCtxt)
        Right (hdrs, records) -> do
            participantMetrics <- lookupParticipantFieldMapping v
            case checkDetermineMapping participantMetrics (getHeaders hdrs) records of
                Right columnCompatibleMetrics -> do
                    liftIO (BS.writeFile ("/tmp/" ++ uploadFileName) uploadFileContents)
                    pure
                        $ okResponse []
                            $ JSON.recordEncoding -- TODO: not record encoding
                                $ JSON.Record vi
                                    $      "csv_upload_id" JSON..= uploadFileName
                                        <> "column_samples" JSON..= extractSampleColumns 5 hdrs records
                                        <> "suggested_mapping" JSON..= headerMappingJSON columnCompatibleMetrics
                Left err -> do
                    liftIO (print ("failed to determine mapping", err))
                    -- if column check failed, then don't save csv file and response is error
                    pure (forbiddenResponse reqCtxt) -- place holder for error

getHeaders :: CSV.Header -> [Text]
getHeaders hdrs =
  (V.toList . fmap TE.decodeUtf8) hdrs

runParticipantUpload :: ActionRoute (Id Volume)
runParticipantUpload = action POST (pathJSON >/> pathId </< "runParticipantUpload") $ \vi -> withAuth $ do
    v <- getVolume PermissionEDIT vi
    reqCtxt <- peek
    (csvUploadId :: String, selectedMapping :: JSON.Value) <- runForm (Nothing) $ do
        csrfForm
        (uploadId :: String) <- "csv_upload_id" .:> deform
        mapping <- "selected_mapping" .:> deform
        pure (uploadId, mapping)
    -- TODO: resolve csv id to absolute path; http error if unknown
    uploadFileContents <- (liftIO . BS.readFile) ("/tmp/" ++ csvUploadId)
    case ATTO.parseOnly (CSVP.csvWithHeader CSVP.defaultDecodeOptions) uploadFileContents of
        Left err ->
            pure (forbiddenResponse reqCtxt) -- TODO: better error
        Right (hdrs, records) -> do
            participantActiveMetrics <- lookupParticipantFieldMapping v
            let eMpngs = JSON.parseEither (parseMapping participantActiveMetrics) selectedMapping
            liftIO $ print ("upload id", csvUploadId, "mapping", eMpngs)
            -- TODO: validate mappings against allowed/detected data types
            let Right mpngs = eMpngs -- TODO: handle either above
            eRes <- (liftIO . runImport records) mpngs
            pure
                $ okResponse []
                    $ JSON.recordEncoding -- TODO: not record encoding
                        $ JSON.Record vi $ "succeeded" JSON..= True

-- TODO: unit tests
-- TODO: validator should also ensure each col mentioned is a real column
parseMapping :: [Metric] -> JSON.Value -> JSON.Parser ParticipantFieldMapping -- TODO: split into simple parser + validator
parseMapping participantActiveMetrics val = do
    (entries :: [HeaderMappingEntry]) <- JSON.parseJSON val
    let metricField :: Map Text Text
        metricField = (MAP.fromList . fmap (\e -> (hmeMetricName e, hmeCsvField e))) entries
    ParticipantFieldMapping
        <$> getFieldIfUsed "id" metricField
        <*> getFieldIfUsed "info" metricField
        <*> getFieldIfUsed "description" metricField
        <*> getFieldIfUsed "birthdate" metricField
        <*> getFieldIfUsed "gender" metricField
        <*> getFieldIfUsed "race" metricField
        <*> getFieldIfUsed "ethnicity" metricField
        <*> getFieldIfUsed "gestationalage" metricField  -- space
        <*> getFieldIfUsed "pregnancyterm" metricField  -- space
        <*> getFieldIfUsed "birthweight" metricField  -- space
        <*> getFieldIfUsed "disability" metricField
        <*> getFieldIfUsed "language" metricField
        <*> getFieldIfUsed "country" metricField
        <*> getFieldIfUsed "state" metricField
        <*> getFieldIfUsed "setting" metricField
  where
    getFieldIfUsed :: Text -> Map Text Text -> JSON.Parser (Maybe Text)
    getFieldIfUsed metricSymbolicName metricField =
        case findMetricBySymbolicName metricSymbolicName of
            Just metric ->
                case MAP.lookup metricSymbolicName metricField of
                    Just csvField -> pure (Just csvField)
                    Nothing -> fail "missing expected participant metric" -- TODO: name metric
            Nothing ->
                pure Nothing
    findMetricBySymbolicName :: Text -> Maybe Metric
    findMetricBySymbolicName symbolicName =
        L.find (\m -> (T.filter (/= ' ') . T.toLower . metricName) m == symbolicName) participantActiveMetrics

runImport :: V.Vector CSV.NamedRecord -> ParticipantFieldMapping -> IO (V.Vector ()) -- TODO: error or count
runImport records mapping =
    mapM
        (\record -> createRecord mapping record)
        records

extractSampleColumns :: Int -> CSV.Header -> V.Vector CSV.NamedRecord -> [JSON.Value]
extractSampleColumns maxSamples hdrs records =
    (V.toList . fmap (\hdr -> sampleColumnJson maxSamples hdr (extractColumn hdr))) hdrs
  where
    extractColumn :: BS.ByteString -> [Maybe BS.ByteString]  -- Should error out if receive nothing
    extractColumn hdr = 
        V.toList (fmap (\rowMap -> HMP.lookup hdr rowMap) records)

sampleColumnJson :: Int -> BS.ByteString -> [Maybe BS.ByteString] -> JSON.Value
sampleColumnJson maxSamples hdr columnValues =
    let
        uniqueSamples = (take maxSamples . L.nub . fmap (maybe "" id)) columnValues
    in
        JSON.object [
            "column_name" JSON..= hdr
          , "samples" JSON..= uniqueSamples
          ]

createRecord :: ParticipantFieldMapping -> CSV.NamedRecord -> IO () -- TODO: error or record
createRecord mapping csvRecord = do
     -- record <- makeRecord
     -- for each field
     case pfmId mapping of --     metricId = getId name
         Just idCol ->
             case HMP.lookup (TE.encodeUtf8 idCol) csvRecord of
                 Just fieldVal ->
                     -- getCSVField row field
                     print ("save measure", "ID", fieldVal)
                 Nothing ->
                     pure () -- TODO: error, impossible?
         Nothing ->
             pure () -- field isn't used by this volume, so don't need to save the measure
     pure ()
