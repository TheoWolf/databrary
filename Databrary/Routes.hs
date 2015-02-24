{-# LANGUAGE OverloadedStrings #-}
module Databrary.Routes
  ( routes
  ) where

import Control.Applicative ((<|>))
import Control.Monad (msum, mfilter, guard)

import qualified Databrary.Web.Route as R
import Databrary.Action
import Databrary.Controller.Root
import Databrary.Controller.Login
import Databrary.Controller.Register
import Databrary.Controller.Token
import Databrary.Controller.Party
import Databrary.Controller.Authorize
import Databrary.Controller.Volume
import Databrary.Controller.Record
import Databrary.Controller.SlotAsset
import Databrary.Controller.Citation
import Databrary.Controller.Angular
import Databrary.Controller.Static

act :: RouteAction q -> R.RouteM (Action q)
act ra = do
  R.final
  _ <- mfilter (actionMethod ra ==) R.method
  return $ routeAction ra

routes :: R.RouteM AppAction
routes = do
  api <- R.route
  let
    json = guard (api == JSON)
    html = guard (api == HTML)
  msum 
    [                                 act (viewRoot api)
    , "login" >>             (html >> act viewLogin)
                                  <|> act (postLogin api)
    , "register" >>          (html >> act viewRegister)
                                  <|> act (postRegister api)
    , "password" >>          (html >> act viewPasswordReset)
                                  <|> act (postPasswordReset api)
    , R.route >>= \t ->               act (viewLoginToken api t)
                                  <|> act (postPasswordToken api t)
    , R.route >>= \p -> msum          -- /party/P
      [                               act (viewParty api p)
      ,                               act (postParty api p)
      , html >> "edit" >>             act (viewPartyForm p)
      , R.route >>= \a ->    (html >> act (viewAuthorize p a))
                                  <|> act (postAuthorize api p a)
      ]
    , "party" >>                      act (createParty api)
                                  <|> act (searchParty api)
    , R.route >>= \v ->               act (viewVolume api v)
    , R.route >>= \c ->
        R.route >>= \a ->
               (html >> "download" >> act (downloadSlotAsset c a))
    , R.route >>= \r ->               act (viewRecord api r)

    , json >> msum                    -- /api
      [ "cite" >>                     act getCitation
      ]
    , html >> msum                    -- /
      [ "public" >> R.route >>=       act . staticPublicFile
      , "constants.js" >>             act angularConstants
      ]
    ]
