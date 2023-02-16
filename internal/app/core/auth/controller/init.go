package controller

import (
	"github.com/labstack/echo/v4"

	"github.com/rodericusifo/echo-template/internal/app/core/auth/service"
)

type AuthController struct {
	AuthService service.IAuthService
}

func InitAuthController(authService service.IAuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (authController *AuthController) Mount(group *echo.Group) {
	group.POST("/login", authController.LoginAuth)
}
