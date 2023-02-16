package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var (
	dBinder    = new(echo.DefaultBinder)
	cValidator = InitCustomValidator()
)

type customValidator struct {
	Validator *validator.Validate
}

func InitCustomValidator() *customValidator {
	return &customValidator{Validator: validator.New()}
}

func (cv *customValidator) Validate(i any) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}
