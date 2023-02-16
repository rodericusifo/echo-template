package controller

import (
	"github.com/labstack/echo/v4"

	"github.com/rodericusifo/echo-template/mocks"
)

var (
	mockApp         *echo.Echo
	mockAuthService *mocks.IAuthService
	authController  *AuthController
)

var (
	mockEmail, mockPassword, mockJWTToken string
)

func SetupTestAuthController() {
	mockApp = echo.New()

	mockAuthService = new(mocks.IAuthService)

	authController = InitAuthController(mockAuthService)
	authController.Mount(mockApp.Group("/auth"))

	mockEmail = "john@gmail.com"
	mockPassword = "john1223"

	mockJWTToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjhlYTc3OGJjLTM5NTgtNGU5Zi04ZmEyLWE4YTlhZDhmMmFiMSIsIm5hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AZ21haWwuY29tIiwicm9sZSI6IkFETUlOIiwiZXhwIjoxNjc3MDc5NzgxfQ.bndXk_BggjadIF2Rwluxc-3tPr-ArfWVYTZ5y03wHU8"
}
