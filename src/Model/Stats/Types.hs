{-# LANGUAGE OverloadedStrings, RecordWildCards #-}
module Model.Stats.Types
  ( SiteStats(..)
  ) where

import qualified Data.Aeson as JSON
import qualified Data.Array.Unboxed as A
import Data.Int (Int64)
import qualified Data.Map.Strict as M

import Model.Permission.Types
import Model.Offset
import Model.Id.Types
import Model.Category.Types

data SiteStats = SiteStats
  { statsAuthorizedSite :: !(A.Array Permission Int64)
  , statsVolumes, statsVolumesShared :: !Int64
  , statsAssets :: !Int64
  , statsAssetDuration :: !Offset
  , statsAssetBytes :: !Int64
  , statsRecords :: !(M.Map (Id Category) Int64)
  }

instance JSON.ToJSON SiteStats where
  toJSON SiteStats{..} = JSON.object
    [ "authorized" JSON..= A.elems statsAuthorizedSite
    , "volumes" JSON..= statsVolumes
    , "shared" JSON..= statsVolumesShared
    , "assets" JSON..= statsAssets
    , "duration" JSON..= statsAssetDuration
    , "bytes" JSON..= statsAssetBytes
    , "records" JSON..= M.mapKeys show statsRecords
    ]
