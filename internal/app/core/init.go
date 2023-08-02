package core

import (
	"github.com/labstack/echo/v4"

	"github.com/rodericusifo/echo-template/internal/pkg/util"

	echojwt "github.com/labstack/echo-jwt/v4"

	auth_controller "github.com/rodericusifo/echo-template/internal/app/core/auth/controller"
	employee_controller "github.com/rodericusifo/echo-template/internal/app/core/employee/controller"
	wire_core_service_auth "github.com/rodericusifo/echo-template/wire/core/service/auth"
	wire_core_service_employee "github.com/rodericusifo/echo-template/wire/core/service/employee"
)

func InitRoutes(app *echo.Echo) {
	{
		auth := app.Group("/auth")
		authService := wire_core_service_auth.AuthService()
		authController := auth_controller.InitAuthController(authService)
		authController.Mount(auth)
	}
	{
		employee := app.Group("/employees")
		employee.Use(echojwt.WithConfig(*util.GetJWTAuthConfig()))
		employeeService := wire_core_service_employee.EmployeeService()
		employeeController := employee_controller.InitEmployeeController(employeeService)
		employeeController.Mount(employee)
	}
}
