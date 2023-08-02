package employee

import (
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/pkg/types"
	"github.com/rodericusifo/echo-template/pkg/util"
)

func (r *PostgresEmployeeDatabaseSQLRepository) CountAllEmployee(query *types.Query) (int, error) {
	count := 0
	employees := make([]*sql.Employee, 0)
	tableName := sql.Employee{}.TableName()

	q := r.db

	querySlice := util.GenerateSQLSelectQuerySlice(
		tableName,
		[]types.SelectOperation{
			{Field: "id"},
		},
	)
	q = q.Select(querySlice)

	if query != nil {
		q = util.BuildQuery(tableName, q, query)
	}

	q = q.Table(tableName).Find(&employees)

	if err := q.Error; err != nil {
		return count, err
	}

	count = int(q.RowsAffected)

	return count, nil
}
