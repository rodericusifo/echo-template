package resource

import (
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/pkg/types"
)

func (r *EmployeeResource) GetListEmployeeAndCount(query *types.Query) ([]*sql.Employee, int, error) {
	return r.EmployeeDatabaseSQLRepository.GetListEmployeeAndCount(query)
}
