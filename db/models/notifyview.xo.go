// Package xo_models contains the types for schema 'public'.
package xo_models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
)

// NotifyView represents a row from '"public"."notify_view"'.
type NotifyView struct {
	Target   sql.NullInt64  `json:"target"`   // target
	Notice   sql.NullInt64  `json:"notice"`   // notice
	Delivery NoticeDelivery `json:"delivery"` // delivery
}
