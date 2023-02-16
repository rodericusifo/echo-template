package request

import (
	"time"
)

type UpdateEmployeeRequestBody struct {
	Address  *string    `json:"address" validate:"omitempty"`
	Age      *int       `json:"age" validate:"omitempty,min=0"`
	Birthday *time.Time `json:"birthday" validate:"omitempty"`
}

func (r *UpdateEmployeeRequestBody) CustomValidateRequestBody() error {
	return nil
}

type UpdateEmployeeRequestParams struct {
	XID string `param:"xid" validate:"required,uuid4"`
}

func (r *UpdateEmployeeRequestParams) CustomValidateRequestParams() error {
	return nil
}
