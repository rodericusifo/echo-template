package user

import (
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
)

type UserSeedPayload struct {
	XID      string            `validate:"required,uuid4"`
	Name     string            `validate:"required"`
	Email    string            `validate:"required,email"`
	Password string            `validate:"required,min=8"`
	Role     constant.UserRole `validate:"required,oneof=ADMIN"`
}

func (r *UserSeedPayload) CustomValidatePayload() error {
	return nil
}
