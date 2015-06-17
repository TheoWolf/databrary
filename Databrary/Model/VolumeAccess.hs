{-# LANGUAGE TemplateHaskell, RecordWildCards, OverloadedStrings #-}
module Databrary.Model.VolumeAccess
  ( module Databrary.Model.VolumeAccess.Types
  , lookupVolumeAccess
  , lookupVolumeAccessParty
  , lookupPartyVolumeAccess
  , changeVolumeAccess
  , volumeAccessProvidesADMIN
  , volumeAccessJSON
  , lookupVolumeActivity
  ) where

import Data.Int (Int64)
import Data.Maybe (catMaybes)

import Databrary.Ops
import Databrary.Has (peek, view)
import qualified Databrary.JSON as JSON
import Databrary.Service.DB
import Databrary.Model.SQL
import Databrary.Model.Time
import Databrary.Model.Id.Types
import Databrary.Model.Permission.Types
import Databrary.Model.Identity.Types
import Databrary.Model.Party.Types
import Databrary.Model.Volume.Types
import Databrary.Model.Audit
import Databrary.Model.VolumeAccess.Types
import Databrary.Model.VolumeAccess.SQL

lookupVolumeAccess :: (MonadDB m, MonadHasIdentity c m) => Volume -> Permission -> m [VolumeAccess]
lookupVolumeAccess vol perm = do
  ident <- peek
  dbQuery $(selectQuery (selectVolumeAccess 'vol 'ident) "$WHERE volume_access.individual >= ${perm} ORDER BY individual DESC, children DESC")

lookupVolumeAccessParty :: (MonadDB m, MonadHasIdentity c m) => Volume -> Id Party -> m (Maybe VolumeAccess)
lookupVolumeAccessParty vol p = do
  ident <- peek
  dbQuery1 $(selectQuery (selectVolumeAccessParty 'vol 'ident) "WHERE party.id = ${p}")

lookupPartyVolumeAccess :: (MonadDB m, MonadHasIdentity c m) => Party -> Permission -> m [VolumeAccess]
lookupPartyVolumeAccess p perm = do
  ident <- peek
  dbQuery $(selectQuery (selectPartyVolumeAccess 'p 'ident) "$WHERE volume_access.individual >= ${perm} ORDER BY individual DESC, children DESC")

changeVolumeAccess :: (MonadAudit c m) => VolumeAccess -> m Bool
changeVolumeAccess va = do
  ident <- getAuditIdentity
  if volumeAccessIndividual va == PermissionNONE
    then dbExecute1 $(deleteVolumeAccess 'ident 'va)
    else (0 <) . fst <$> updateOrInsert
      $(updateVolumeAccess 'ident 'va)
      $(insertVolumeAccess 'ident 'va)

volumeAccessProvidesADMIN :: VolumeAccess -> Bool
volumeAccessProvidesADMIN VolumeAccess{ volumeAccessChildren   = PermissionADMIN, volumeAccessParty = p } = accessMember     p == PermissionADMIN
volumeAccessProvidesADMIN VolumeAccess{ volumeAccessIndividual = PermissionADMIN, volumeAccessParty = p } = accessPermission p == PermissionADMIN
volumeAccessProvidesADMIN _ = False

volumeAccessJSON :: VolumeAccess -> JSON.Object
volumeAccessJSON VolumeAccess{..} = JSON.object $ catMaybes
  [ ("individual" JSON..= volumeAccessIndividual) <? (volumeAccessIndividual >= PermissionNONE)
  , ("children"   JSON..= volumeAccessChildren)   <? (volumeAccessChildren   >= PermissionNONE)
  ]

lookupVolumeActivity :: (MonadDB m, MonadHasIdentity c m) => Int -> m [(Timestamp, Volume)]
lookupVolumeActivity limit = do
  ident :: Identity <- peek
  dbQuery $(selectQuery (selectVolumeActivity 'ident) "$WHERE audit.audit_action = 'add' AND audit.party = 0 AND audit.children > 'NONE' ORDER BY audit.audit_time DESC LIMIT ${fromIntegral limit :: Int64}")