package resource

import (
	"github.com/rodericusifo/echo-template/pkg/types"
)

func (r *EmployeeResource) CountAllEmployee(query *types.Query) (int, error) {
	return r.EmployeeDatabaseSQLRepository.CountAllEmployee(query)
}
