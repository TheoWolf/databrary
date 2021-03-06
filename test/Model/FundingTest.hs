{-# LANGUAGE OverloadedStrings, ScopedTypeVariables, GeneralizedNewtypeDeriving
   , TypeSynonymInstances, MultiParamTypeClasses, FlexibleInstances, RecordWildCards #-}
module Model.FundingTest where

import Test.Tasty.HUnit

import Model.Funding
import Model.TypeOrphans ()
import TestHarness

-- session driving a variety of functions in the module
unit_Funding_examples :: Assertion
unit_Funding_examples =
    withinTestTransaction (\cn -> do
        let ctxt = TestContext { ctxConn = Just cn }
            fndr = Funder (Id 1) "The Big Fund"
        matches <-
            runReaderT
              (do
                  _ <- addFunder fndr
                  findFunders "Big")
              ctxt
        matches @?= [fndr])
