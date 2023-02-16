package resource

import (
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/pkg/types"
)

func (r *EmployeeResource) GetEmployee(query *types.Query) (*sql.Employee, error) {
	return r.EmployeeDatabaseSQLRepository.GetEmployee(query)
}
