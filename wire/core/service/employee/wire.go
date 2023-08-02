//go:build wireinject
// +build wireinject

package employee

import (
	"github.com/google/wire"

	internal_pkg_util "github.com/rodericusifo/echo-template/internal/pkg/util"

	employee_resource "github.com/rodericusifo/echo-template/internal/app/core/employee/resource"
	employee_service "github.com/rodericusifo/echo-template/internal/app/core/employee/service"
	employee_database_repository "github.com/rodericusifo/echo-template/internal/app/repository/database/sql/employee"
)

func EmployeeService() employee_service.IEmployeeService {
	wire.Build(
		internal_pkg_util.GetPostgresDBConnection,
		employee_database_repository.InitPostgresEmployeeDatabaseSQLRepository,
		employee_resource.InitEmployeeResource,
		employee_service.InitEmployeeService,
	)
	return &employee_service.EmployeeService{}
}
