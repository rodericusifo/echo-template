package request

import (
	"time"
)

type CreateEmployeeRequestBody struct {
	Name     string     `json:"name" validate:"required"`
	Email    string     `json:"email" validate:"required,email"`
	Address  *string    `json:"address" validate:"omitempty"`
	Age      *int       `json:"age" validate:"omitempty,min=0"`
	Birthday *time.Time `json:"birthday" validate:"omitempty"`
}

func (r *CreateEmployeeRequestBody) CustomValidateRequestBody() error {
	return nil
}
