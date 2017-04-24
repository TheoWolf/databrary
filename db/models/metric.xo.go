// Package xo_models contains the types for schema 'public'.
package xo_models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Metric represents a row from '"public"."metric"'.
type Metric struct {
	ID          int64            `json:"id"`          // id
	Category    int16            `json:"category"`    // category
	Name        string           `json:"name"`        // name
	Release     Release          `json:"release"`     // release
	Type        DataType         `json:"type"`        // type
	Options     []sql.NullString `json:"options"`     // options
	Assumed     sql.NullString   `json:"assumed"`     // assumed
	Description sql.NullString   `json:"description"` // description
	Required    sql.NullBool     `json:"required"`    // required

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Metric exists in the database.
func (m *Metric) Exists() bool {
	return m._exists
}

// Deleted provides information if the Metric has been deleted from the database.
func (m *Metric) Deleted() bool {
	return m._deleted
}

// Insert inserts the Metric to the database.
func (m *Metric) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if m._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO "public"."metric" (` +
		`"category", "name", "release", "type", "options", "assumed", "description", "required"` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) RETURNING "id"`

	// run query
	XOLog(sqlstr, m.Category, m.Name, m.Release, m.Type, m.Options, m.Assumed, m.Description, m.Required)
	err = db.QueryRow(sqlstr, m.Category, m.Name, m.Release, m.Type, m.Options, m.Assumed, m.Description, m.Required).Scan(&m.ID)
	if err != nil {
		return err
	}

	// set existence
	m._exists = true

	return nil
}

// Update updates the Metric in the database.
func (m *Metric) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !m._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if m._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE "public"."metric" SET (` +
		`"category", "name", "release", "type", "options", "assumed", "description", "required"` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) WHERE "id" = $9`

	// run query
	XOLog(sqlstr, m.Category, m.Name, m.Release, m.Type, m.Options, m.Assumed, m.Description, m.Required, m.ID)
	_, err = db.Exec(sqlstr, m.Category, m.Name, m.Release, m.Type, m.Options, m.Assumed, m.Description, m.Required, m.ID)
	return err
}

// Save saves the Metric to the database.
func (m *Metric) Save(db XODB) error {
	if m.Exists() {
		return m.Update(db)
	}

	return m.Insert(db)
}

// Upsert performs an upsert for Metric.
//
// NOTE: PostgreSQL 9.5+ only
func (m *Metric) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if m._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO "public"."metric" (` +
		`"id", "category", "name", "release", "type", "options", "assumed", "description", "required"` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`) ON CONFLICT ("id") DO UPDATE SET (` +
		`"id", "category", "name", "release", "type", "options", "assumed", "description", "required"` +
		`) = (` +
		`EXCLUDED."id", EXCLUDED."category", EXCLUDED."name", EXCLUDED."release", EXCLUDED."type", EXCLUDED."options", EXCLUDED."assumed", EXCLUDED."description", EXCLUDED."required"` +
		`)`

	// run query
	XOLog(sqlstr, m.ID, m.Category, m.Name, m.Release, m.Type, m.Options, m.Assumed, m.Description, m.Required)
	_, err = db.Exec(sqlstr, m.ID, m.Category, m.Name, m.Release, m.Type, m.Options, m.Assumed, m.Description, m.Required)
	if err != nil {
		return err
	}

	// set existence
	m._exists = true

	return nil
}

// Delete deletes the Metric from the database.
func (m *Metric) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !m._exists {
		return nil
	}

	// if deleted, bail
	if m._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM "public"."metric" WHERE "id" = $1`

	// run query
	XOLog(sqlstr, m.ID)
	_, err = db.Exec(sqlstr, m.ID)
	if err != nil {
		return err
	}

	// set deleted
	m._deleted = true

	return nil
}

// Category returns the Category associated with the Metric's Category (category).
//
// Generated from foreign key 'metric_category_fkey'.
func (m *Metric) CategoryByCategoryID(db XODB) (*Category, error) {
	return CategoryByID(db, m.Category)
}

// MetricByCategoryName retrieves a row from '"public"."metric"' as a Metric.
//
// Generated from index 'metric_category_name_key'.
func MetricByCategoryName(db XODB, category int16, name string) (*Metric, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`"id", "category", "name", "release", "type", "options", "assumed", "description", "required" ` +
		`FROM "public"."metric" ` +
		`WHERE "category" = $1 AND "name" = $2`

	// run query
	XOLog(sqlstr, category, name)
	m := Metric{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, category, name).Scan(&m.ID, &m.Category, &m.Name, &m.Release, &m.Type, &m.Options, &m.Assumed, &m.Description, &m.Required)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// MetricByID retrieves a row from '"public"."metric"' as a Metric.
//
// Generated from index 'metric_pkey'.
func MetricByID(db XODB, id int64) (*Metric, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`"id", "category", "name", "release", "type", "options", "assumed", "description", "required" ` +
		`FROM "public"."metric" ` +
		`WHERE "id" = $1`

	// run query
	XOLog(sqlstr, id)
	m := Metric{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&m.ID, &m.Category, &m.Name, &m.Release, &m.Type, &m.Options, &m.Assumed, &m.Description, &m.Required)
	if err != nil {
		return nil, err
	}

	return &m, nil
}
