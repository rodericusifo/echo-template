package core

import (
	"github.com/labstack/echo/v4"

	"github.com/rodericusifo/echo-template/internal/pkg/util"

	echojwt "github.com/labstack/echo-jwt/v4"

	wire_core_controller_auth "github.com/rodericusifo/echo-template/wire/core/controller/auth"
	wire_core_controller_employee "github.com/rodericusifo/echo-template/wire/core/controller/employee"
)

func InitRoutes(app *echo.Echo) {
	{
		auth := app.Group("/auth")
		authController := wire_core_controller_auth.AuthController()
		authController.Mount(auth)
	}
	{
		employee := app.Group("/employees")
		employee.Use(echojwt.WithConfig(*util.GetJWTAuthConfig()))
		employeeController := wire_core_controller_employee.EmployeeController()
		employeeController.Mount(employee)
	}
}
