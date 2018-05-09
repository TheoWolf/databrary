{-# LANGUAGE TemplateHaskell, QuasiQuotes, RecordWildCards, OverloadedStrings, DataKinds #-}
module Databrary.Model.Tag
  ( module Databrary.Model.Tag.Types
  , lookupTag
  , lookupTags
  , findTags
  , addTag
  , lookupVolumeTagUseRows
  , addTagUse
  , removeTagUse
  , lookupTopTagWeight
  , lookupTagCoverage
  , lookupSlotTagCoverage
  , lookupSlotKeywords
  , tagWeightJSON
  , tagCoverageJSON
  ) where

import Control.Applicative (empty, pure)
import Control.Monad (guard)
import qualified Data.ByteString.Char8 as BSC
import Data.Int (Int64)
import Data.Maybe (fromMaybe)
import Data.Monoid ((<>))
import qualified Data.String
-- import Database.PostgreSQL.Typed (pgSQL)
import Database.PostgreSQL.Typed.Types

import Databrary.Has (peek)
import qualified Databrary.JSON as JSON
import Databrary.Service.DB
import Databrary.Model.SQL
import Databrary.Model.Party.Types
import Databrary.Model.Identity.Types
import Databrary.Model.Volume.Types
import Databrary.Model.Container.Types
import Databrary.Model.Slot.Types
import Databrary.Model.Tag.Types
import Databrary.Model.Tag.SQL

lookupTag :: MonadDB c m => TagName -> m (Maybe Tag)
lookupTag n =
  dbQuery1 $(selectQuery selectTag "$WHERE tag.name = ${n}::varchar")

lookupTags :: MonadDB c m => m [Tag]
lookupTags = do
  let _tenv_a6Dq8 = unknownPGTypeEnv
  rows <- dbQuery -- (selectQuery selectTag "")
    (mapQuery2
                      (BSC.concat
                         [Data.String.fromString "SELECT tag.id,tag.name FROM tag "])
              (\ [_cid_a6Dq9, _cname_a6Dqa]
                 -> (Database.PostgreSQL.Typed.Types.pgDecodeColumnNotNull
                       _tenv_a6Dq8
                       (Database.PostgreSQL.Typed.Types.PGTypeProxy ::
                          Database.PostgreSQL.Typed.Types.PGTypeName "integer")
                       _cid_a6Dq9, 
                     Database.PostgreSQL.Typed.Types.pgDecodeColumnNotNull
                       _tenv_a6Dq8
                       (Database.PostgreSQL.Typed.Types.PGTypeProxy ::
                          Database.PostgreSQL.Typed.Types.PGTypeName "character varying")
                       _cname_a6Dqa)))
  pure
    (fmap
      (\ (vid_a6Dpn, vname_a6Dpo) -> Tag vid_a6Dpn vname_a6Dpo)
      rows)

findTags :: MonadDB c m => TagName -> Int -> m [Tag]
findTags (TagName n) lim = -- TagName restrictions obviate pattern escaping
  dbQuery $(selectQuery selectTag "$WHERE tag.name LIKE ${n `BSC.snoc` '%'}::varchar LIMIT ${fromIntegral lim :: Int64}")

addTag :: MonadDB c m => TagName -> m Tag
addTag n = do
  let _tenv_a6GtM = unknownPGTypeEnv
  row <- dbQuery1' -- [pgSQL|!SELECT get_tag(${n})|]
    (mapQuery2
      ((\ _p_a6GtN ->
                    (BSC.concat
                       [Data.String.fromString "SELECT get_tag(",
                        Database.PostgreSQL.Typed.Types.pgEscapeParameter
                          _tenv_a6GtM
                          (Database.PostgreSQL.Typed.Types.PGTypeProxy ::
                             Database.PostgreSQL.Typed.Types.PGTypeName "character varying")
                          _p_a6GtN,
                        Data.String.fromString ")"]))
      n)
      (\ [_cget_tag_a6GtO]
               -> (Database.PostgreSQL.Typed.Types.pgDecodeColumnNotNull
                     _tenv_a6GtM
                     (Database.PostgreSQL.Typed.Types.PGTypeProxy ::
                        Database.PostgreSQL.Typed.Types.PGTypeName "integer")
                     _cget_tag_a6GtO)))
  pure ((`Tag` n) row)

lookupVolumeTagUseRows :: MonadDB c m => Volume -> m [TagUseRow]
lookupVolumeTagUseRows v =
  dbQuery $(selectQuery selectTagUseRow "JOIN container ON tag_use.container = container.id WHERE container.volume = ${volumeId $ volumeRow v} ORDER BY container.id")

addTagUse :: MonadDB c m => TagUse -> m Bool
addTagUse t = either (const False) id <$> do
  dbTryJust (guard . isExclusionViolation)
    $ dbExecute1 (if tagKeyword t
      then $(insertTagUse True 't)
      else $(insertTagUse False 't))

removeTagUse :: MonadDB c m => TagUse -> m Int
removeTagUse t =
  dbExecute
    (if tagKeyword t
      then $(deleteTagUse True 't)
      else $(deleteTagUse False 't))

lookupTopTagWeight :: MonadDB c m => Int -> m [TagWeight]
lookupTopTagWeight lim =
  dbQuery $(selectQuery (selectTagWeight "") "$!ORDER BY weight DESC LIMIT ${fromIntegral lim :: Int64}")

emptyTagCoverage :: Tag -> Container -> TagCoverage
emptyTagCoverage t c = TagCoverage (TagWeight t 0) c [] [] []

lookupTagCoverage :: (MonadDB c m, MonadHasIdentity c m) => Tag -> Slot -> m TagCoverage
lookupTagCoverage t (Slot c s) = do
  ident <- peek
  fromMaybe (emptyTagCoverage t c) <$> dbQuery1 (($ c) . ($ t) <$> $(selectQuery (selectTagCoverage 'ident "WHERE container = ${containerId $ containerRow c} AND segment && ${s} AND tag = ${tagId t}") "$!"))

lookupSlotTagCoverage :: (MonadDB c m, MonadHasIdentity c m) => Slot -> Int -> m [TagCoverage]
lookupSlotTagCoverage slot lim = do
  ident <- peek
  dbQuery $(selectQuery (selectSlotTagCoverage 'ident 'slot) "$!ORDER BY weight DESC LIMIT ${fromIntegral lim :: Int64}")

lookupSlotKeywords :: (MonadDB c m) => Slot -> m [Tag]
lookupSlotKeywords Slot{..} =
  dbQuery $(selectQuery selectTag "JOIN keyword_use ON id = tag WHERE container = ${containerId $ containerRow slotContainer} AND segment = ${slotSegment}")

tagWeightJSON :: JSON.ToObject o => TagWeight -> JSON.Record TagName o
tagWeightJSON TagWeight{..} = JSON.Record (tagName tagWeightTag) $
  "weight" JSON..= tagWeightWeight

tagCoverageJSON :: JSON.ToObject o => TagCoverage -> JSON.Record TagName o
tagCoverageJSON TagCoverage{..} = tagWeightJSON tagCoverageWeight `JSON.foldObjectIntoRec`
 (   "coverage" JSON..= tagCoverageSegments
  <> "keyword" `JSON.kvObjectOrEmpty` (if null tagCoverageKeywords then empty else pure tagCoverageKeywords)
  <> "vote"    `JSON.kvObjectOrEmpty` (if null tagCoverageVotes then empty else pure tagCoverageVotes))