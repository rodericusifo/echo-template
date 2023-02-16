package request

import (
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
)

type RequestUser struct {
	ID    uint              `validate:"required"`
	XID   string            `validate:"required,uuid4"`
	Name  string            `validate:"required"`
	Email string            `validate:"required,email"`
	Role  constant.UserRole `validate:"required,oneof=ADMIN NON_ADMIN"`
}

func (r *RequestUser) CustomValidateRequestUser() error {
	return nil
}
