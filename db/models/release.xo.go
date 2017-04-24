// Package xo_models contains the types for schema 'public'.
package xo_models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql/driver"
	"errors"
)

// Release is the 'release' enum type from schema 'public'.
type Release uint16

const (
	// ReleasePrivate is the 'PRIVATE' Release.
	ReleasePrivate = Release(1)

	// ReleaseShared is the 'SHARED' Release.
	ReleaseShared = Release(2)

	// ReleaseExcerpts is the 'EXCERPTS' Release.
	ReleaseExcerpts = Release(3)

	// ReleasePublic is the 'PUBLIC' Release.
	ReleasePublic = Release(4)
)

// String returns the string value of the Release.
func (r Release) String() string {
	var enumVal string

	switch r {
	case ReleasePrivate:
		enumVal = "PRIVATE"

	case ReleaseShared:
		enumVal = "SHARED"

	case ReleaseExcerpts:
		enumVal = "EXCERPTS"

	case ReleasePublic:
		enumVal = "PUBLIC"
	}

	return enumVal
}

// MarshalText marshals Release into text.
func (r Release) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

// UnmarshalText unmarshals Release from text.
func (r *Release) UnmarshalText(text []byte) error {
	switch string(text) {
	case "PRIVATE":
		*r = ReleasePrivate

	case "SHARED":
		*r = ReleaseShared

	case "EXCERPTS":
		*r = ReleaseExcerpts

	case "PUBLIC":
		*r = ReleasePublic

	default:
		return errors.New("invalid Release")
	}

	return nil
}

// Value satisfies the sql/driver.Valuer interface for Release.
func (r Release) Value() (driver.Value, error) {
	return r.String(), nil
}

// Scan satisfies the database/sql.Scanner interface for Release.
func (r *Release) Scan(src interface{}) error {
	buf, ok := src.([]byte)
	if !ok {
		return errors.New("invalid Release")
	}

	return r.UnmarshalText(buf)
}
