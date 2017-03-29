package models

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/databrary/databrary/models/netaddr"
	"time"
)

type (
	// Logs of all activities on the site, including access and modifications to any data.
	// Each table has an associated audit table inheriting from this one.
	Audit struct {
		Time      time.Time    `json:"audit_time" db:"audit_time"`
		UserID    uint         `json:"audit_user" db:"audit_user"` // references party
		IpAddress netaddr.Inet `json:"audit_ip" db:"audit_ip"`
		Action    Action       `json:"audit_action" db:"audit_action"`
	}
)

// Action Enum
type Action string

// The various activities for which we keep audit records (in audit or a derived table).
const (
	ActionATTEMPT   Action = Action("ATTEMPT")
	ActionOPEN      Action = Action("OPEN")
	ActionCLOSE     Action = Action("CLOSE")
	ActionADD       Action = Action("ADD")
	ActionCHANGE    Action = Action("CHANGE")
	ActionREMOVE    Action = Action("REMOVE")
	ActionSUPERUSER Action = Action("SUPERUSER")
)

func (act Action) Value() (driver.Value, error) {
	// value needs to be a base driver.Value type
	// such as bool.
	return string(act), nil
}

func (act *Action) Scan(value interface{}) error {
	if value == nil {
		return ActionErrDb{
			msg: "got nil value from database for Action",
		}
	}
	if exposure_val, err := driver.String.ConvertValue(value); err == nil {
		if v, ok := exposure_val.(Action); ok {
			*act = Action(v)
			return nil
		}
	}
	return ActionErrScn{
		msg: "failed to scan Action",
	}
}

// house keeping

// checking to make sure interface is implemented
var _ sql.Scanner = (*Action)(nil)
var _ driver.Valuer = ActionSUPERUSER

type ActionErr interface {
	error
}

type ActionErrDb struct {
	msg string
}

func (e ActionErrDb) Error() string {
	return fmt.Sprintf("%s", e.msg)
}

type ActionErrScn struct {
	msg string
}

func (e ActionErrScn) Error() string {
	return fmt.Sprintln("%s", e.msg)
}

// consult volume.Test_error