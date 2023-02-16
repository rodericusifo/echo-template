package request

type LoginAuthRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *LoginAuthRequestBody) CustomValidateRequestBody() error {
	return nil
}
