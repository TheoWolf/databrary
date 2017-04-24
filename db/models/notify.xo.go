// Package xo_models contains the types for schema 'public'.
package xo_models

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
)

// Notify represents a row from '"public"."notify"'.
type Notify struct {
	Target   int64          `json:"target"`   // target
	Notice   int16          `json:"notice"`   // notice
	Delivery NoticeDelivery `json:"delivery"` // delivery

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Notify exists in the database.
func (n *Notify) Exists() bool {
	return n._exists
}

// Deleted provides information if the Notify has been deleted from the database.
func (n *Notify) Deleted() bool {
	return n._deleted
}

// Insert inserts the Notify to the database.
func (n *Notify) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if n._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO "public"."notify" (` +
		`"target", "notice", "delivery"` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)`

	// run query
	XOLog(sqlstr, n.Target, n.Notice, n.Delivery)
	err = db.QueryRow(sqlstr, n.Target, n.Notice, n.Delivery).Scan(&n.Notice)
	if err != nil {
		return err
	}

	// set existence
	n._exists = true

	return nil
}

// Update updates the Notify in the database.
func (n *Notify) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !n._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if n._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE "public"."notify" SET (` +
		`"target", "delivery"` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE "notice" = $3`

	// run query
	XOLog(sqlstr, n.Target, n.Delivery, n.Notice)
	_, err = db.Exec(sqlstr, n.Target, n.Delivery, n.Notice)
	return err
}

// Save saves the Notify to the database.
func (n *Notify) Save(db XODB) error {
	if n.Exists() {
		return n.Update(db)
	}

	return n.Insert(db)
}

// Upsert performs an upsert for Notify.
//
// NOTE: PostgreSQL 9.5+ only
func (n *Notify) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if n._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO "public"."notify" (` +
		`"target", "notice", "delivery"` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT ("notice") DO UPDATE SET (` +
		`"target", "notice", "delivery"` +
		`) = (` +
		`EXCLUDED."target", EXCLUDED."notice", EXCLUDED."delivery"` +
		`)`

	// run query
	XOLog(sqlstr, n.Target, n.Notice, n.Delivery)
	_, err = db.Exec(sqlstr, n.Target, n.Notice, n.Delivery)
	if err != nil {
		return err
	}

	// set existence
	n._exists = true

	return nil
}

// Delete deletes the Notify from the database.
func (n *Notify) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !n._exists {
		return nil
	}

	// if deleted, bail
	if n._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM "public"."notify" WHERE "notice" = $1`

	// run query
	XOLog(sqlstr, n.Notice)
	_, err = db.Exec(sqlstr, n.Notice)
	if err != nil {
		return err
	}

	// set deleted
	n._deleted = true

	return nil
}

// Notice returns the Notice associated with the Notify's Notice (notice).
//
// Generated from foreign key 'notify_notice_fkey'.
func (n *Notify) NoticeByNoticeID(db XODB) (*Notice, error) {
	return NoticeByID(db, n.Notice)
}

// Account returns the Account associated with the Notify's Target (target).
//
// Generated from foreign key 'notify_target_fkey'.
func (n *Notify) Account(db XODB) (*Account, error) {
	return AccountByID(db, n.Target)
}

// NotifyByTargetNotice retrieves a row from '"public"."notify"' as a Notify.
//
// Generated from index 'notify_pkey'.
func NotifyByTargetNotice(db XODB, target int64, notice int16) (*Notify, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`"target", "notice", "delivery" ` +
		`FROM "public"."notify" ` +
		`WHERE "target" = $1 AND "notice" = $2`

	// run query
	XOLog(sqlstr, target, notice)
	n := Notify{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, target, notice).Scan(&n.Target, &n.Notice, &n.Delivery)
	if err != nil {
		return nil, err
	}

	return &n, nil
}
