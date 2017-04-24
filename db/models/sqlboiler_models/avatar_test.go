// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testAvatars(t *testing.T) {
	t.Parallel()

	query := Avatars(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAvatarsDelete(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	avatar := &Avatar{}
	if err = randomize.Struct(seed, avatar, avatarDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = avatar.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = avatar.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Avatars(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAvatarsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	avatar := &Avatar{}
	if err = randomize.Struct(seed, avatar, avatarDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = avatar.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Avatars(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Avatars(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAvatarsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	avatar := &Avatar{}
	if err = randomize.Struct(seed, avatar, avatarDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = avatar.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AvatarSlice{avatar}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Avatars(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAvatarsExists(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	avatar := &Avatar{}
	if err = randomize.Struct(seed, avatar, avatarDBTypes, true, avatarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = avatar.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := AvatarExists(tx, avatar.Party)
	if err != nil {
		t.Errorf("Unable to check if Avatar exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AvatarExistsG to return true, but got false.")
	}
}
func testAvatarsFind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	avatar := &Avatar{}
	if err = randomize.Struct(seed, avatar, avatarDBTypes, true, avatarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = avatar.Insert(tx); err != nil {
		t.Error(err)
	}

	avatarFound, err := FindAvatar(tx, avatar.Party)
	if err != nil {
		t.Error(err)
	}

	if avatarFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAvatarsBind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	avatar := &Avatar{}
	if err = randomize.Struct(seed, avatar, avatarDBTypes, true, avatarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = avatar.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Avatars(tx).Bind(avatar); err != nil {
		t.Error(err)
	}
}

func testAvatarsOne(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	avatar := &Avatar{}
	if err = randomize.Struct(seed, avatar, avatarDBTypes, true, avatarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = avatar.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Avatars(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAvatarsAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	avatarOne := &Avatar{}
	avatarTwo := &Avatar{}
	if err = randomize.Struct(seed, avatarOne, avatarDBTypes, false, avatarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}
	if err = randomize.Struct(seed, avatarTwo, avatarDBTypes, false, avatarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = avatarOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = avatarTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Avatars(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAvatarsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	avatarOne := &Avatar{}
	avatarTwo := &Avatar{}
	if err = randomize.Struct(seed, avatarOne, avatarDBTypes, false, avatarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}
	if err = randomize.Struct(seed, avatarTwo, avatarDBTypes, false, avatarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = avatarOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = avatarTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Avatars(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func avatarBeforeInsertHook(e boil.Executor, o *Avatar) error {
	*o = Avatar{}
	return nil
}

func avatarAfterInsertHook(e boil.Executor, o *Avatar) error {
	*o = Avatar{}
	return nil
}

func avatarAfterSelectHook(e boil.Executor, o *Avatar) error {
	*o = Avatar{}
	return nil
}

func avatarBeforeUpdateHook(e boil.Executor, o *Avatar) error {
	*o = Avatar{}
	return nil
}

func avatarAfterUpdateHook(e boil.Executor, o *Avatar) error {
	*o = Avatar{}
	return nil
}

func avatarBeforeDeleteHook(e boil.Executor, o *Avatar) error {
	*o = Avatar{}
	return nil
}

func avatarAfterDeleteHook(e boil.Executor, o *Avatar) error {
	*o = Avatar{}
	return nil
}

func avatarBeforeUpsertHook(e boil.Executor, o *Avatar) error {
	*o = Avatar{}
	return nil
}

func avatarAfterUpsertHook(e boil.Executor, o *Avatar) error {
	*o = Avatar{}
	return nil
}

func testAvatarsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Avatar{}
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	o := &Avatar{}
	if err = randomize.Struct(seed, o, avatarDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Avatar object: %s", err)
	}

	AddAvatarHook(boil.BeforeInsertHook, avatarBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	avatarBeforeInsertHooks = []AvatarHook{}

	AddAvatarHook(boil.AfterInsertHook, avatarAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	avatarAfterInsertHooks = []AvatarHook{}

	AddAvatarHook(boil.AfterSelectHook, avatarAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	avatarAfterSelectHooks = []AvatarHook{}

	AddAvatarHook(boil.BeforeUpdateHook, avatarBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	avatarBeforeUpdateHooks = []AvatarHook{}

	AddAvatarHook(boil.AfterUpdateHook, avatarAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	avatarAfterUpdateHooks = []AvatarHook{}

	AddAvatarHook(boil.BeforeDeleteHook, avatarBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	avatarBeforeDeleteHooks = []AvatarHook{}

	AddAvatarHook(boil.AfterDeleteHook, avatarAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	avatarAfterDeleteHooks = []AvatarHook{}

	AddAvatarHook(boil.BeforeUpsertHook, avatarBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	avatarBeforeUpsertHooks = []AvatarHook{}

	AddAvatarHook(boil.AfterUpsertHook, avatarAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	avatarAfterUpsertHooks = []AvatarHook{}
}
func testAvatarsInsert(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	avatar := &Avatar{}
	if err = randomize.Struct(seed, avatar, avatarDBTypes, true, avatarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = avatar.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Avatars(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAvatarsInsertWhitelist(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	avatar := &Avatar{}
	if err = randomize.Struct(seed, avatar, avatarDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = avatar.Insert(tx, avatarColumns...); err != nil {
		t.Error(err)
	}

	count, err := Avatars(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAvatarToOneAssetUsingAsset(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Avatar
	var foreign Asset

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, avatarDBTypes, true, avatarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, assetDBTypes, true, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Asset = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.AssetByFk(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := AvatarSlice{&local}
	if err = local.L.LoadAsset(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Asset == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Asset = nil
	if err = local.L.LoadAsset(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Asset == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAvatarToOnePartyUsingParty(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Avatar
	var foreign Party

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, avatarDBTypes, true, avatarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, partyDBTypes, true, partyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Party struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Party = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.PartyByFk(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := AvatarSlice{&local}
	if err = local.L.LoadParty(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Party == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Party = nil
	if err = local.L.LoadParty(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Party == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAvatarToOneSetOpAssetUsingAsset(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Avatar
	var b, c Asset

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, avatarDBTypes, false, strmangle.SetComplement(avatarPrimaryKeyColumns, avatarColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, assetDBTypes, false, strmangle.SetComplement(assetPrimaryKeyColumns, assetColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, assetDBTypes, false, strmangle.SetComplement(assetPrimaryKeyColumns, assetColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Asset{&b, &c} {
		err = a.SetAsset(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Asset != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Avatars[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Asset != x.ID {
			t.Error("foreign key was wrong value", a.Asset)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Asset))
		reflect.Indirect(reflect.ValueOf(&a.Asset)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.Asset != x.ID {
			t.Error("foreign key was wrong value", a.Asset, x.ID)
		}
	}
}
func testAvatarToOneSetOpPartyUsingParty(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Avatar
	var b, c Party

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, avatarDBTypes, false, strmangle.SetComplement(avatarPrimaryKeyColumns, avatarColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, partyDBTypes, false, strmangle.SetComplement(partyPrimaryKeyColumns, partyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, partyDBTypes, false, strmangle.SetComplement(partyPrimaryKeyColumns, partyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Party{&b, &c} {
		err = a.SetParty(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Party != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Avatar != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Party != x.ID {
			t.Error("foreign key was wrong value", a.Party)
		}

		if exists, err := AvatarExists(tx, a.Party); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testAvatarsReload(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	avatar := &Avatar{}
	if err = randomize.Struct(seed, avatar, avatarDBTypes, true, avatarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = avatar.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = avatar.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAvatarsReloadAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	avatar := &Avatar{}
	if err = randomize.Struct(seed, avatar, avatarDBTypes, true, avatarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = avatar.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AvatarSlice{avatar}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAvatarsSelect(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	avatar := &Avatar{}
	if err = randomize.Struct(seed, avatar, avatarDBTypes, true, avatarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = avatar.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Avatars(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	avatarDBTypes = map[string]string{`Asset`: `integer`, `Party`: `integer`}
	_             = bytes.MinRead
)

func testAvatarsUpdate(t *testing.T) {
	t.Parallel()

	if len(avatarColumns) == len(avatarPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	avatar := &Avatar{}
	if err = randomize.Struct(seed, avatar, avatarDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = avatar.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Avatars(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, avatar, avatarDBTypes, true, avatarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	if err = avatar.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAvatarsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(avatarColumns) == len(avatarPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	avatar := &Avatar{}
	if err = randomize.Struct(seed, avatar, avatarDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = avatar.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Avatars(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, avatar, avatarDBTypes, true, avatarPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(avatarColumns, avatarPrimaryKeyColumns) {
		fields = avatarColumns
	} else {
		fields = strmangle.SetComplement(
			avatarColumns,
			avatarPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(avatar))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := AvatarSlice{avatar}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAvatarsUpsert(t *testing.T) {
	t.Parallel()

	if len(avatarColumns) == len(avatarPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed
	// Attempt the INSERT side of an UPSERT

	avatar := Avatar{}
	if err = randomize.Struct(seed, &avatar, avatarDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = avatar.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Avatar: %s", err)
	}

	count, err := Avatars(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT

	if err = randomize.Struct(seed, &avatar, avatarDBTypes, false, avatarPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Avatar struct: %s", err)
	}

	if err = avatar.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Avatar: %s", err)
	}

	count, err = Avatars(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
