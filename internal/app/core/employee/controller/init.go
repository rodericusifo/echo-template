package controller

import (
	"github.com/labstack/echo/v4"

	"github.com/rodericusifo/echo-template/internal/app/core/employee/service"
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/internal/pkg/middleware"
)

type EmployeeController struct {
	EmployeeService service.IEmployeeService
}

func InitEmployeeController(employeeService service.IEmployeeService) *EmployeeController {
	return &EmployeeController{EmployeeService: employeeService}
}

func (employeeController *EmployeeController) Mount(group *echo.Group) {
	group.POST("/create", employeeController.CreateEmployee, middleware.UserRequest(), middleware.UserRolesPermission(constant.ADMIN))
	group.GET("/list", employeeController.GetListEmployee, middleware.UserRequest(), middleware.UserRolesPermission(constant.ADMIN))
	group.GET("/:xid/detail", employeeController.GetEmployee, middleware.UserRequest(), middleware.UserRolesPermission(constant.ADMIN))
	group.PATCH("/:xid/update", employeeController.UpdateEmployee, middleware.UserRequest(), middleware.UserRolesPermission(constant.ADMIN))
	group.DELETE("/:xid/delete", employeeController.DeleteEmployee, middleware.UserRequest(), middleware.UserRolesPermission(constant.ADMIN))
}
