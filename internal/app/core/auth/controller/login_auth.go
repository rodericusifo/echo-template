package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/rodericusifo/echo-template/internal/app/core/auth/controller/request"
	"github.com/rodericusifo/echo-template/internal/app/core/auth/controller/response"
	"github.com/rodericusifo/echo-template/internal/app/core/auth/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/pkg/validator"

	pkg_response "github.com/rodericusifo/echo-template/pkg/response"
)

func (c *AuthController) LoginAuth(ctx echo.Context) error {
	reqBody := new(request.LoginAuthRequestBody)
	if err := validator.ValidateRequestBody(ctx, reqBody); err != nil {
		return err
	}

	authLoginDtoRes, err := c.AuthService.LoginAuth(&input.LoginAuthDTO{
		Email:    reqBody.Email,
		Password: reqBody.Password,
	})
	if err != nil {
		return err
	}

	loginAuthRes := &response.LoginAuthResponse{
		Token: authLoginDtoRes.Token,
	}

	return ctx.JSON(http.StatusOK, pkg_response.ResponseSuccess("auth login success", loginAuthRes, nil))
}
