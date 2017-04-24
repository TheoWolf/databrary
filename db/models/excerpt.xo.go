// Package xo_models contains the types for schema 'public'.
package xo_models

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"

	"github.com/databrary/databrary/db/models/custom_types"
)

// Excerpt represents a row from '"public"."excerpt"'.
type Excerpt struct {
	Asset   int64                `json:"asset"`   // asset
	Segment custom_types.Segment `json:"segment"` // segment
	Release Release              `json:"release"` // release

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Excerpt exists in the database.
func (e *Excerpt) Exists() bool {
	return e._exists
}

// Deleted provides information if the Excerpt has been deleted from the database.
func (e *Excerpt) Deleted() bool {
	return e._deleted
}

// Insert inserts the Excerpt to the database.
func (e *Excerpt) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO "public"."excerpt" (` +
		`"asset", "segment", "release"` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)`

	// run query
	XOLog(sqlstr, e.Asset, e.Segment, e.Release)
	err = db.QueryRow(sqlstr, e.Asset, e.Segment, e.Release).Scan(&e.Segment)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Update updates the Excerpt in the database.
func (e *Excerpt) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !e._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if e._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE "public"."excerpt" SET (` +
		`"asset", "release"` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE "segment" = $3`

	// run query
	XOLog(sqlstr, e.Asset, e.Release, e.Segment)
	_, err = db.Exec(sqlstr, e.Asset, e.Release, e.Segment)
	return err
}

// Save saves the Excerpt to the database.
func (e *Excerpt) Save(db XODB) error {
	if e.Exists() {
		return e.Update(db)
	}

	return e.Insert(db)
}

// Upsert performs an upsert for Excerpt.
//
// NOTE: PostgreSQL 9.5+ only
func (e *Excerpt) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO "public"."excerpt" (` +
		`"asset", "segment", "release"` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT ("segment") DO UPDATE SET (` +
		`"asset", "segment", "release"` +
		`) = (` +
		`EXCLUDED."asset", EXCLUDED."segment", EXCLUDED."release"` +
		`)`

	// run query
	XOLog(sqlstr, e.Asset, e.Segment, e.Release)
	_, err = db.Exec(sqlstr, e.Asset, e.Segment, e.Release)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Delete deletes the Excerpt from the database.
func (e *Excerpt) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !e._exists {
		return nil
	}

	// if deleted, bail
	if e._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM "public"."excerpt" WHERE "segment" = $1`

	// run query
	XOLog(sqlstr, e.Segment)
	_, err = db.Exec(sqlstr, e.Segment)
	if err != nil {
		return err
	}

	// set deleted
	e._deleted = true

	return nil
}

// SlotAsset returns the SlotAsset associated with the Excerpt's Asset (asset).
//
// Generated from foreign key 'excerpt_asset_fkey'.
func (e *Excerpt) SlotAsset(db XODB) (*SlotAsset, error) {
	return SlotAssetByAsset(db, e.Asset)
}

// ExcerptByAssetSegment retrieves a row from '"public"."excerpt"' as a Excerpt.
//
// Generated from index 'excerpt_pkey'.
func ExcerptByAssetSegment(db XODB, asset int64, segment custom_types.Segment) (*Excerpt, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`"asset", "segment", "release" ` +
		`FROM "public"."excerpt" ` +
		`WHERE "asset" = $1 AND "segment" = $2`

	// run query
	XOLog(sqlstr, asset, segment)
	e := Excerpt{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, asset, segment).Scan(&e.Asset, &e.Segment, &e.Release)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
