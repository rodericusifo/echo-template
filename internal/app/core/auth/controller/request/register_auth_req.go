package request

import (
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
)

type RegisterAuthRequestBody struct {
	Name     string            `validate:"required"`
	Email    string            `validate:"required,email"`
	Password string            `validate:"required,min=8"`
	Role     constant.UserRole `validate:"required,oneof=NON_ADMIN"`
}

func (r *RegisterAuthRequestBody) CustomValidateRequestBody() error {
	return nil
}
