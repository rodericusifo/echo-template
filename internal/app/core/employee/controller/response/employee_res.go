package response

import (
	"time"
)

type EmployeeResponse struct {
	XID       string     `json:"xid,omitempty"`
	Name      string     `json:"name,omitempty"`
	Email     string     `json:"email,omitempty"`
	Address   *string    `json:"address,omitempty"`
	Age       *int       `json:"age,omitempty"`
	Birthday  *time.Time `json:"birthday,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
