package resource

import (
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/internal/app/repository/database/sql/employee"
	"github.com/rodericusifo/echo-template/pkg/types"
)

type IEmployeeResource interface {
	CreateEmployee(payload *sql.Employee) error
	UpdateEmployee(payload *sql.Employee) error
	DeleteEmployee(payload *sql.Employee) error
	GetListEmployeeAndCount(query *types.Query) ([]*sql.Employee, int, error)
	GetEmployee(query *types.Query) (*sql.Employee, error)
	CountAllEmployee(query *types.Query) (int, error)
}

type EmployeeResource struct {
	EmployeeDatabaseSQLRepository employee.IEmployeeDatabaseSQLRepository
}

func InitEmployeeResource(employeeDatabaseSQLRepository employee.IEmployeeDatabaseSQLRepository) IEmployeeResource {
	return &EmployeeResource{
		EmployeeDatabaseSQLRepository: employeeDatabaseSQLRepository,
	}
}
