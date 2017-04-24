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

func testFormats(t *testing.T) {
	t.Parallel()

	query := Formats(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFormatsDelete(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	format := &Format{}
	if err = randomize.Struct(seed, format, formatDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = format.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = format.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Formats(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFormatsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	format := &Format{}
	if err = randomize.Struct(seed, format, formatDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = format.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Formats(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Formats(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFormatsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	format := &Format{}
	if err = randomize.Struct(seed, format, formatDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = format.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FormatSlice{format}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Formats(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFormatsExists(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	format := &Format{}
	if err = randomize.Struct(seed, format, formatDBTypes, true, formatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = format.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FormatExists(tx, format.ID)
	if err != nil {
		t.Errorf("Unable to check if Format exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FormatExistsG to return true, but got false.")
	}
}
func testFormatsFind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	format := &Format{}
	if err = randomize.Struct(seed, format, formatDBTypes, true, formatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = format.Insert(tx); err != nil {
		t.Error(err)
	}

	formatFound, err := FindFormat(tx, format.ID)
	if err != nil {
		t.Error(err)
	}

	if formatFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFormatsBind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	format := &Format{}
	if err = randomize.Struct(seed, format, formatDBTypes, true, formatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = format.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Formats(tx).Bind(format); err != nil {
		t.Error(err)
	}
}

func testFormatsOne(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	format := &Format{}
	if err = randomize.Struct(seed, format, formatDBTypes, true, formatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = format.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Formats(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFormatsAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	formatOne := &Format{}
	formatTwo := &Format{}
	if err = randomize.Struct(seed, formatOne, formatDBTypes, false, formatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}
	if err = randomize.Struct(seed, formatTwo, formatDBTypes, false, formatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = formatOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = formatTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Formats(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFormatsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	formatOne := &Format{}
	formatTwo := &Format{}
	if err = randomize.Struct(seed, formatOne, formatDBTypes, false, formatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}
	if err = randomize.Struct(seed, formatTwo, formatDBTypes, false, formatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = formatOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = formatTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Formats(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func formatBeforeInsertHook(e boil.Executor, o *Format) error {
	*o = Format{}
	return nil
}

func formatAfterInsertHook(e boil.Executor, o *Format) error {
	*o = Format{}
	return nil
}

func formatAfterSelectHook(e boil.Executor, o *Format) error {
	*o = Format{}
	return nil
}

func formatBeforeUpdateHook(e boil.Executor, o *Format) error {
	*o = Format{}
	return nil
}

func formatAfterUpdateHook(e boil.Executor, o *Format) error {
	*o = Format{}
	return nil
}

func formatBeforeDeleteHook(e boil.Executor, o *Format) error {
	*o = Format{}
	return nil
}

func formatAfterDeleteHook(e boil.Executor, o *Format) error {
	*o = Format{}
	return nil
}

func formatBeforeUpsertHook(e boil.Executor, o *Format) error {
	*o = Format{}
	return nil
}

func formatAfterUpsertHook(e boil.Executor, o *Format) error {
	*o = Format{}
	return nil
}

func testFormatsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Format{}
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	o := &Format{}
	if err = randomize.Struct(seed, o, formatDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Format object: %s", err)
	}

	AddFormatHook(boil.BeforeInsertHook, formatBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	formatBeforeInsertHooks = []FormatHook{}

	AddFormatHook(boil.AfterInsertHook, formatAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	formatAfterInsertHooks = []FormatHook{}

	AddFormatHook(boil.AfterSelectHook, formatAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	formatAfterSelectHooks = []FormatHook{}

	AddFormatHook(boil.BeforeUpdateHook, formatBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	formatBeforeUpdateHooks = []FormatHook{}

	AddFormatHook(boil.AfterUpdateHook, formatAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	formatAfterUpdateHooks = []FormatHook{}

	AddFormatHook(boil.BeforeDeleteHook, formatBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	formatBeforeDeleteHooks = []FormatHook{}

	AddFormatHook(boil.AfterDeleteHook, formatAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	formatAfterDeleteHooks = []FormatHook{}

	AddFormatHook(boil.BeforeUpsertHook, formatBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	formatBeforeUpsertHooks = []FormatHook{}

	AddFormatHook(boil.AfterUpsertHook, formatAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	formatAfterUpsertHooks = []FormatHook{}
}
func testFormatsInsert(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	format := &Format{}
	if err = randomize.Struct(seed, format, formatDBTypes, true, formatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = format.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Formats(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFormatsInsertWhitelist(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	format := &Format{}
	if err = randomize.Struct(seed, format, formatDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = format.Insert(tx, formatColumns...); err != nil {
		t.Error(err)
	}

	count, err := Formats(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFormatToManyAssets(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Format
	var b, c Asset

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, formatDBTypes, true, formatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, assetDBTypes, false, assetColumnsWithDefault...)
	randomize.Struct(seed, &c, assetDBTypes, false, assetColumnsWithDefault...)

	b.Format = a.ID
	c.Format = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	asset, err := a.AssetsByFk(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range asset {
		if v.Format == b.Format {
			bFound = true
		}
		if v.Format == c.Format {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := FormatSlice{&a}
	if err = a.L.LoadAssets(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Assets); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Assets = nil
	if err = a.L.LoadAssets(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Assets); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", asset)
	}
}

func testFormatToManyAddOpAssets(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Format
	var b, c, d, e Asset

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, formatDBTypes, false, strmangle.SetComplement(formatPrimaryKeyColumns, formatColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Asset{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, assetDBTypes, false, strmangle.SetComplement(assetPrimaryKeyColumns, assetColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Asset{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAssets(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.Format {
			t.Error("foreign key was wrong value", a.ID, first.Format)
		}
		if a.ID != second.Format {
			t.Error("foreign key was wrong value", a.ID, second.Format)
		}

		if first.R.Format != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Format != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Assets[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Assets[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AssetsByFk(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testFormatsReload(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	format := &Format{}
	if err = randomize.Struct(seed, format, formatDBTypes, true, formatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = format.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = format.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFormatsReloadAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	format := &Format{}
	if err = randomize.Struct(seed, format, formatDBTypes, true, formatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = format.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FormatSlice{format}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFormatsSelect(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	format := &Format{}
	if err = randomize.Struct(seed, format, formatDBTypes, true, formatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = format.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Formats(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	formatDBTypes = map[string]string{`Extension`: `ARRAYcharacter varying`, `ID`: `smallint`, `Mimetype`: `character varying`, `Name`: `text`}
	_             = bytes.MinRead
)

func testFormatsUpdate(t *testing.T) {
	t.Parallel()

	if len(formatColumns) == len(formatPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	format := &Format{}
	if err = randomize.Struct(seed, format, formatDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = format.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Formats(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, format, formatDBTypes, true, formatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	if err = format.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFormatsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(formatColumns) == len(formatPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed

	format := &Format{}
	if err = randomize.Struct(seed, format, formatDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = format.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Formats(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, format, formatDBTypes, true, formatPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(formatColumns, formatPrimaryKeyColumns) {
		fields = formatColumns
	} else {
		fields = strmangle.SetComplement(
			formatColumns,
			formatPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(format))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FormatSlice{format}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFormatsUpsert(t *testing.T) {
	t.Parallel()

	if len(formatColumns) == len(formatPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	// this is a hack because if randomize isn't used compiler will complain
	// but if seed isn't then compiler will complain too
	_ = seed
	// Attempt the INSERT side of an UPSERT

	format := Format{}
	if err = randomize.Struct(seed, &format, formatDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = format.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Format: %s", err)
	}

	count, err := Formats(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT

	if err = randomize.Struct(seed, &format, formatDBTypes, false, formatPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Format struct: %s", err)
	}

	if err = format.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Format: %s", err)
	}

	count, err = Formats(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
