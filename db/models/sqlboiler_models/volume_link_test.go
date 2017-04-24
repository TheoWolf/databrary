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

func testVolumeLinks(t *testing.T) {
	t.Parallel()

	query := VolumeLinks(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testVolumeLinksDelete(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = volumeLink.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testVolumeLinksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = VolumeLinks(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testVolumeLinksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := VolumeLinkSlice{volumeLink}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testVolumeLinksExists(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true, volumeLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := VolumeLinkExists(tx, volumeLink.Volume, volumeLink.URL)
	if err != nil {
		t.Errorf("Unable to check if VolumeLink exists: %s", err)
	}
	if !e {
		t.Errorf("Expected VolumeLinkExistsG to return true, but got false.")
	}
}
func testVolumeLinksFind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true, volumeLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	volumeLinkFound, err := FindVolumeLink(tx, volumeLink.Volume, volumeLink.URL)
	if err != nil {
		t.Error(err)
	}

	if volumeLinkFound == nil {
		t.Error("want a record, got nil")
	}
}
func testVolumeLinksBind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true, volumeLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = VolumeLinks(tx).Bind(volumeLink); err != nil {
		t.Error(err)
	}
}

func testVolumeLinksOne(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true, volumeLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := VolumeLinks(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testVolumeLinksAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	volumeLinkOne := &VolumeLink{}
	volumeLinkTwo := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLinkOne, volumeLinkDBTypes, false, volumeLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}
	if err = randomize.Struct(seed, volumeLinkTwo, volumeLinkDBTypes, false, volumeLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLinkOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = volumeLinkTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := VolumeLinks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testVolumeLinksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	volumeLinkOne := &VolumeLink{}
	volumeLinkTwo := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLinkOne, volumeLinkDBTypes, false, volumeLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}
	if err = randomize.Struct(seed, volumeLinkTwo, volumeLinkDBTypes, false, volumeLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLinkOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = volumeLinkTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func volumeLinkBeforeInsertHook(e boil.Executor, o *VolumeLink) error {
	*o = VolumeLink{}
	return nil
}

func volumeLinkAfterInsertHook(e boil.Executor, o *VolumeLink) error {
	*o = VolumeLink{}
	return nil
}

func volumeLinkAfterSelectHook(e boil.Executor, o *VolumeLink) error {
	*o = VolumeLink{}
	return nil
}

func volumeLinkBeforeUpdateHook(e boil.Executor, o *VolumeLink) error {
	*o = VolumeLink{}
	return nil
}

func volumeLinkAfterUpdateHook(e boil.Executor, o *VolumeLink) error {
	*o = VolumeLink{}
	return nil
}

func volumeLinkBeforeDeleteHook(e boil.Executor, o *VolumeLink) error {
	*o = VolumeLink{}
	return nil
}

func volumeLinkAfterDeleteHook(e boil.Executor, o *VolumeLink) error {
	*o = VolumeLink{}
	return nil
}

func volumeLinkBeforeUpsertHook(e boil.Executor, o *VolumeLink) error {
	*o = VolumeLink{}
	return nil
}

func volumeLinkAfterUpsertHook(e boil.Executor, o *VolumeLink) error {
	*o = VolumeLink{}
	return nil
}

func testVolumeLinksHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &VolumeLink{}
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	o := &VolumeLink{}
	if err = randomize.Struct(seed, o, volumeLinkDBTypes, false); err != nil {
		t.Errorf("Unable to randomize VolumeLink object: %s", err)
	}

	AddVolumeLinkHook(boil.BeforeInsertHook, volumeLinkBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	volumeLinkBeforeInsertHooks = []VolumeLinkHook{}

	AddVolumeLinkHook(boil.AfterInsertHook, volumeLinkAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	volumeLinkAfterInsertHooks = []VolumeLinkHook{}

	AddVolumeLinkHook(boil.AfterSelectHook, volumeLinkAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	volumeLinkAfterSelectHooks = []VolumeLinkHook{}

	AddVolumeLinkHook(boil.BeforeUpdateHook, volumeLinkBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	volumeLinkBeforeUpdateHooks = []VolumeLinkHook{}

	AddVolumeLinkHook(boil.AfterUpdateHook, volumeLinkAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	volumeLinkAfterUpdateHooks = []VolumeLinkHook{}

	AddVolumeLinkHook(boil.BeforeDeleteHook, volumeLinkBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	volumeLinkBeforeDeleteHooks = []VolumeLinkHook{}

	AddVolumeLinkHook(boil.AfterDeleteHook, volumeLinkAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	volumeLinkAfterDeleteHooks = []VolumeLinkHook{}

	AddVolumeLinkHook(boil.BeforeUpsertHook, volumeLinkBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	volumeLinkBeforeUpsertHooks = []VolumeLinkHook{}

	AddVolumeLinkHook(boil.AfterUpsertHook, volumeLinkAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	volumeLinkAfterUpsertHooks = []VolumeLinkHook{}
}
func testVolumeLinksInsert(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true, volumeLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testVolumeLinksInsertWhitelist(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx, volumeLinkColumns...); err != nil {
		t.Error(err)
	}

	count, err := VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testVolumeLinkToOneVolumeUsingVolume(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local VolumeLink
	var foreign Volume

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, volumeLinkDBTypes, true, volumeLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, volumeDBTypes, true, volumeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Volume struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Volume = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.VolumeByFk(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := VolumeLinkSlice{&local}
	if err = local.L.LoadVolume(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Volume == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Volume = nil
	if err = local.L.LoadVolume(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Volume == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testVolumeLinkToOneSetOpVolumeUsingVolume(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a VolumeLink
	var b, c Volume

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, volumeLinkDBTypes, false, strmangle.SetComplement(volumeLinkPrimaryKeyColumns, volumeLinkColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, volumeDBTypes, false, strmangle.SetComplement(volumePrimaryKeyColumns, volumeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, volumeDBTypes, false, strmangle.SetComplement(volumePrimaryKeyColumns, volumeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Volume{&b, &c} {
		err = a.SetVolume(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Volume != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.VolumeLinks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Volume != x.ID {
			t.Error("foreign key was wrong value", a.Volume)
		}

		if exists, err := VolumeLinkExists(tx, a.Volume, a.URL); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testVolumeLinksReload(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true, volumeLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = volumeLink.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testVolumeLinksReloadAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true, volumeLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := VolumeLinkSlice{volumeLink}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testVolumeLinksSelect(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true, volumeLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := VolumeLinks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	volumeLinkDBTypes = map[string]string{`Head`: `text`, `URL`: `text`, `Volume`: `integer`}
	_                 = bytes.MinRead
)

func testVolumeLinksUpdate(t *testing.T) {
	t.Parallel()

	if len(volumeLinkColumns) == len(volumeLinkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true, volumeLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	if err = volumeLink.Update(tx); err != nil {
		t.Error(err)
	}
}

func testVolumeLinksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(volumeLinkColumns) == len(volumeLinkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true, volumeLinkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(volumeLinkColumns, volumeLinkPrimaryKeyColumns) {
		fields = volumeLinkColumns
	} else {
		fields = strmangle.SetComplement(
			volumeLinkColumns,
			volumeLinkPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(volumeLink))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := VolumeLinkSlice{volumeLink}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testVolumeLinksUpsert(t *testing.T) {
	t.Parallel()

	if len(volumeLinkColumns) == len(volumeLinkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed
	// Attempt the INSERT side of an UPSERT

	volumeLink := VolumeLink{}
	if err = randomize.Struct(seed, &volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert VolumeLink: %s", err)
	}

	count, err := VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT

	if err = randomize.Struct(seed, &volumeLink, volumeLinkDBTypes, false, volumeLinkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	if err = volumeLink.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert VolumeLink: %s", err)
	}

	count, err = VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
