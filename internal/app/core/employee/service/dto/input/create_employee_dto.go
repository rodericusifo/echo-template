package input

import (
	"time"
)

type CreateEmployeeDTO struct {
	Name     string
	Email    string
	Address  *string
	Age      *int
	Birthday *time.Time
	UserID   uint
}
