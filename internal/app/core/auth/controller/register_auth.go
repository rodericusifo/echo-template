package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/rodericusifo/echo-template/internal/app/core/auth/controller/request"
	"github.com/rodericusifo/echo-template/internal/app/core/auth/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/pkg/validator"

	pkg_response "github.com/rodericusifo/echo-template/pkg/response"
)

func (c *AuthController) RegisterAuth(ctx echo.Context) error {
	reqBody := new(request.RegisterAuthRequestBody)
	if err := validator.ValidateRequestBody(ctx, reqBody); err != nil {
		return err
	}

	err := c.AuthService.RegisterAuth(&input.RegisterAuthDTO{
		Name:     reqBody.Name,
		Email:    reqBody.Email,
		Password: reqBody.Password,
		Role:     reqBody.Role,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, pkg_response.ResponseSuccess[any]("auth register success", nil, nil))
}
