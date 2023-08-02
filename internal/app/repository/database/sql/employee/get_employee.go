package employee

import (
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/pkg/types"
	"github.com/rodericusifo/echo-template/pkg/util"
)

func (r *PostgresEmployeeDatabaseSQLRepository) GetEmployee(query *types.Query) (*sql.Employee, error) {
	employee := new(sql.Employee)
	tableName := sql.Employee{}.TableName()

	q := r.db

	if query != nil {
		q = util.BuildQuery(tableName, q, query)
	}

	if err := q.Table(tableName).First(employee).Error; err != nil {
		return nil, err
	}

	return employee, nil
}
