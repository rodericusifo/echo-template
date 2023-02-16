package output

import (
	"time"
)

type EmployeeDTO struct {
	XID       string
	Name      string
	Email     string
	Address   *string
	Age       *int
	Birthday  *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
