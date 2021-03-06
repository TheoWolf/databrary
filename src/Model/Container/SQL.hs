{-# LANGUAGE TemplateHaskell, OverloadedStrings #-}
module Model.Container.SQL
  ( selectContainerRow
  , selectVolumeContainer
  , selectContainer
  -- for expanded query
  , setContainerId
  ) where

import qualified Language.Haskell.TH as TH

import Model.SQL.Select
import Model.Id.Types
import Model.Volume.SQL
import Model.Release.SQL
import Model.Container.Types

selectContainerRow :: Selector -- ^ @'ContainerRow'@
selectContainerRow = selectColumns 'ContainerRow "container" ["id", "top", "name", "date"]

selectVolumeContainer :: Selector -- ^ @'Volume' -> 'Container'@
selectVolumeContainer = selectJoin 'Container
  [ selectContainerRow
  , maybeJoinOn "container.id = slot_release.container AND slot_release.segment = '(,)'"
    releaseRow
  ]

selectContainer :: TH.Name -- ^ @'Identity'@
  -> Selector -- ^ @'Container'@
selectContainer ident = selectJoin '($)
  [ selectVolumeContainer
  , joinOn "container.volume = volume.id" $ selectVolume ident
  ]

setContainerId :: Container -> Id Container -> Container
setContainerId o i = o{ containerRow = (containerRow o) { containerId = i } }
