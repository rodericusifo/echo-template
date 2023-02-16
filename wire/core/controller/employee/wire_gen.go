// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package employee

import (
	"github.com/rodericusifo/echo-template/internal/app/core/employee/controller"
	"github.com/rodericusifo/echo-template/internal/app/core/employee/resource"
	"github.com/rodericusifo/echo-template/internal/app/core/employee/service"
	"github.com/rodericusifo/echo-template/internal/app/repository/database/sql/employee"
	"github.com/rodericusifo/echo-template/internal/pkg/util"
)

// Injectors from wire.go:

func EmployeeController() *controller.EmployeeController {
	postgresDBSQLConnection := util.GetPostgresDBConnection()
	iEmployeeDatabaseSQLRepository := employee.InitPostgresEmployeeDatabaseSQLRepository(postgresDBSQLConnection)
	iEmployeeResource := resource.InitEmployeeResource(iEmployeeDatabaseSQLRepository)
	iEmployeeService := service.InitEmployeeService(iEmployeeResource)
	employeeController := controller.InitEmployeeController(iEmployeeService)
	return employeeController
}
