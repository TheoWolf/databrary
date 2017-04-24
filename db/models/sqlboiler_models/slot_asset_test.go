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

func testSlotAssets(t *testing.T) {
	t.Parallel()

	query := SlotAssets(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testSlotAssetsDelete(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	slotAsset := SlotAssetRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = slotAsset.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSlotAssetsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	slotAsset := SlotAssetRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = SlotAssets(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSlotAssetsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	slotAsset := SlotAssetRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := SlotAssetSlice{slotAsset}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testSlotAssetsExists(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	slotAsset := SlotAssetRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := SlotAssetExists(tx, slotAsset.Asset)
	if err != nil {
		t.Errorf("Unable to check if SlotAsset exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SlotAssetExistsG to return true, but got false.")
	}
}
func testSlotAssetsFind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	slotAsset := SlotAssetRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	slotAssetFound, err := FindSlotAsset(tx, slotAsset.Asset)
	if err != nil {
		t.Error(err)
	}

	if slotAssetFound == nil {
		t.Error("want a record, got nil")
	}
}
func testSlotAssetsBind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	slotAsset := SlotAssetRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = SlotAssets(tx).Bind(slotAsset); err != nil {
		t.Error(err)
	}
}

func testSlotAssetsOne(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	slotAsset := SlotAssetRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := SlotAssets(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSlotAssetsAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	slotAssetOne := SlotAssetRandom()
	slotAssetTwo := SlotAssetRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAssetOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = slotAssetTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := SlotAssets(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSlotAssetsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	slotAssetOne := SlotAssetRandom()
	slotAssetTwo := SlotAssetRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAssetOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = slotAssetTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func slotAssetBeforeInsertHook(e boil.Executor, o *SlotAsset) error {
	*o = SlotAsset{}
	return nil
}

func slotAssetAfterInsertHook(e boil.Executor, o *SlotAsset) error {
	*o = SlotAsset{}
	return nil
}

func slotAssetAfterSelectHook(e boil.Executor, o *SlotAsset) error {
	*o = SlotAsset{}
	return nil
}

func slotAssetBeforeUpdateHook(e boil.Executor, o *SlotAsset) error {
	*o = SlotAsset{}
	return nil
}

func slotAssetAfterUpdateHook(e boil.Executor, o *SlotAsset) error {
	*o = SlotAsset{}
	return nil
}

func slotAssetBeforeDeleteHook(e boil.Executor, o *SlotAsset) error {
	*o = SlotAsset{}
	return nil
}

func slotAssetAfterDeleteHook(e boil.Executor, o *SlotAsset) error {
	*o = SlotAsset{}
	return nil
}

func slotAssetBeforeUpsertHook(e boil.Executor, o *SlotAsset) error {
	*o = SlotAsset{}
	return nil
}

func slotAssetAfterUpsertHook(e boil.Executor, o *SlotAsset) error {
	*o = SlotAsset{}
	return nil
}

func testSlotAssetsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &SlotAsset{}
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	o := SlotAssetRandom()

	AddSlotAssetHook(boil.BeforeInsertHook, slotAssetBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	slotAssetBeforeInsertHooks = []SlotAssetHook{}

	AddSlotAssetHook(boil.AfterInsertHook, slotAssetAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	slotAssetAfterInsertHooks = []SlotAssetHook{}

	AddSlotAssetHook(boil.AfterSelectHook, slotAssetAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	slotAssetAfterSelectHooks = []SlotAssetHook{}

	AddSlotAssetHook(boil.BeforeUpdateHook, slotAssetBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	slotAssetBeforeUpdateHooks = []SlotAssetHook{}

	AddSlotAssetHook(boil.AfterUpdateHook, slotAssetAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	slotAssetAfterUpdateHooks = []SlotAssetHook{}

	AddSlotAssetHook(boil.BeforeDeleteHook, slotAssetBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	slotAssetBeforeDeleteHooks = []SlotAssetHook{}

	AddSlotAssetHook(boil.AfterDeleteHook, slotAssetAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	slotAssetAfterDeleteHooks = []SlotAssetHook{}

	AddSlotAssetHook(boil.BeforeUpsertHook, slotAssetBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	slotAssetBeforeUpsertHooks = []SlotAssetHook{}

	AddSlotAssetHook(boil.AfterUpsertHook, slotAssetAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	slotAssetAfterUpsertHooks = []SlotAssetHook{}
}
func testSlotAssetsInsert(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	slotAsset := SlotAssetRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSlotAssetsInsertWhitelist(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	slotAsset := SlotAssetRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx, slotAssetColumns...); err != nil {
		t.Error(err)
	}

	count, err := SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSlotAssetToManyAssetExcerpts(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a SlotAsset
	var b, c Excerpt

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, slotAssetDBTypes, true, slotAssetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, excerptDBTypes, false, excerptColumnsWithDefault...)
	randomize.Struct(seed, &c, excerptDBTypes, false, excerptColumnsWithDefault...)

	b.Asset = a.Asset
	c.Asset = a.Asset
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	excerpt, err := a.AssetExcerptsByFk(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range excerpt {
		if v.Asset == b.Asset {
			bFound = true
		}
		if v.Asset == c.Asset {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := SlotAssetSlice{&a}
	if err = a.L.LoadAssetExcerpts(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AssetExcerpts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AssetExcerpts = nil
	if err = a.L.LoadAssetExcerpts(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AssetExcerpts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", excerpt)
	}
}

func testSlotAssetToManyAddOpAssetExcerpts(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a SlotAsset
	var b, c, d, e Excerpt

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, slotAssetDBTypes, false, strmangle.SetComplement(slotAssetPrimaryKeyColumns, slotAssetColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Excerpt{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, excerptDBTypes, false, strmangle.SetComplement(excerptPrimaryKeyColumns, excerptColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Excerpt{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAssetExcerpts(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.Asset != first.Asset {
			t.Error("foreign key was wrong value", a.Asset, first.Asset)
		}
		if a.Asset != second.Asset {
			t.Error("foreign key was wrong value", a.Asset, second.Asset)
		}

		if first.R.Asset != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Asset != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AssetExcerpts[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AssetExcerpts[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AssetExcerptsByFk(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testSlotAssetToOneContainerUsingContainer(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local SlotAsset
	var foreign Container

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, slotAssetDBTypes, true, slotAssetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, containerDBTypes, true, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Container = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.ContainerByFk(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SlotAssetSlice{&local}
	if err = local.L.LoadContainer(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Container == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Container = nil
	if err = local.L.LoadContainer(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Container == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSlotAssetToOneAssetUsingAsset(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local SlotAsset
	var foreign Asset

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, slotAssetDBTypes, true, slotAssetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
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

	slice := SlotAssetSlice{&local}
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

func testSlotAssetToOneSetOpContainerUsingContainer(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a SlotAsset
	var b, c Container

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, slotAssetDBTypes, false, strmangle.SetComplement(slotAssetPrimaryKeyColumns, slotAssetColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, containerDBTypes, false, strmangle.SetComplement(containerPrimaryKeyColumns, containerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, containerDBTypes, false, strmangle.SetComplement(containerPrimaryKeyColumns, containerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Container{&b, &c} {
		err = a.SetContainer(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Container != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SlotAssets[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Container != x.ID {
			t.Error("foreign key was wrong value", a.Container)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Container))
		reflect.Indirect(reflect.ValueOf(&a.Container)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.Container != x.ID {
			t.Error("foreign key was wrong value", a.Container, x.ID)
		}
	}
}
func testSlotAssetToOneSetOpAssetUsingAsset(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a SlotAsset
	var b, c Asset

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, slotAssetDBTypes, false, strmangle.SetComplement(slotAssetPrimaryKeyColumns, slotAssetColumnsWithoutDefault)...); err != nil {
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

		if x.R.SlotAsset != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Asset != x.ID {
			t.Error("foreign key was wrong value", a.Asset)
		}

		if exists, err := SlotAssetExists(tx, a.Asset); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testSlotAssetsReload(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	slotAsset := SlotAssetRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = slotAsset.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testSlotAssetsReloadAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	slotAsset := SlotAssetRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := SlotAssetSlice{slotAsset}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testSlotAssetsSelect(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	slotAsset := SlotAssetRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := SlotAssets(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	slotAssetDBTypes = map[string]string{`Asset`: `integer`, `Container`: `integer`, `Segment`: `USER-DEFINED`}
	_                = bytes.MinRead
)

func testSlotAssetsUpdate(t *testing.T) {
	t.Parallel()

	if len(slotAssetColumns) == len(slotAssetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	slotAsset := SlotAssetRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	slotAsset = SlotAssetRandom()

	if err = slotAsset.Update(tx); err != nil {
		t.Error(err)
	}
}

func testSlotAssetsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(slotAssetColumns) == len(slotAssetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	slotAsset := SlotAssetRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	slotAsset = SlotAssetRandom()

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(slotAssetColumns, slotAssetPrimaryKeyColumns) {
		fields = slotAssetColumns
	} else {
		fields = strmangle.SetComplement(
			slotAssetColumns,
			slotAssetPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(slotAsset))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := SlotAssetSlice{slotAsset}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testSlotAssetsUpsert(t *testing.T) {
	t.Parallel()

	if len(slotAssetColumns) == len(slotAssetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed
	// Attempt the INSERT side of an UPSERT

	slotAsset := SlotAssetRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert SlotAsset: %s", err)
	}

	count, err := SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT

	slotAsset = SlotAssetRandom()

	if err = slotAsset.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert SlotAsset: %s", err)
	}

	count, err = SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
