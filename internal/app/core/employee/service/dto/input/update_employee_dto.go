package input

import (
	"time"
)

type UpdateEmployeeDTO struct {
	XID      string
	Address  *string
	Age      *int
	Birthday *time.Time
	UserID   uint
}
