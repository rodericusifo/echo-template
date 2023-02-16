//go:build wireinject
// +build wireinject

package employee

import (
	"github.com/google/wire"

	internal_pkg_util "github.com/rodericusifo/echo-template/internal/pkg/util"

	employee_controller "github.com/rodericusifo/echo-template/internal/app/core/employee/controller"
	employee_resource "github.com/rodericusifo/echo-template/internal/app/core/employee/resource"
	employee_service "github.com/rodericusifo/echo-template/internal/app/core/employee/service"
	employee_database_repository "github.com/rodericusifo/echo-template/internal/app/repository/database/sql/employee"
)

func EmployeeController() *employee_controller.EmployeeController {
	wire.Build(
		internal_pkg_util.GetPostgresDBConnection,
		employee_database_repository.InitPostgresEmployeeDatabaseSQLRepository,
		employee_resource.InitEmployeeResource,
		employee_service.InitEmployeeService,
		employee_controller.InitEmployeeController,
	)
	return &employee_controller.EmployeeController{}
}
