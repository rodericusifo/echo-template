package employee

import (
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/pkg/types"
	"github.com/rodericusifo/echo-template/pkg/util"
)

func (r *PostgresEmployeeDatabaseSQLRepository) GetListEmployeeAndCount(query *types.Query) ([]*sql.Employee, int, error) {
	count := 0
	employees := make([]*sql.Employee, 0)
	tableName := sql.Employee{}.TableName()

	q := r.db.Table(tableName)

	if query != nil {
		if len(query.Selects) > 0 {
			querySlice := util.GenerateSQLSelectQuerySlice(tableName, util.MergeSlices(true, query.Selects, constant.DEFAULT_SELECT_COLUMNS))
			if query.Distinct {
				q = q.Distinct(querySlice)
			} else {
				q = q.Select(querySlice)
			}
		}
		if len(query.Searches) > 0 {
			queryString, bindValues := util.GenerateSQLWhereQueryStringAndBindValues(tableName, query.Searches)
			q = q.Where(queryString, bindValues...)
		}
		if len(query.Joins) > 0 {
			for _, join := range query.Joins {
				if len(join.Selects) > 0 || len(join.Searches) > 0 {
					qj := r.db
					if len(join.Selects) > 0 {
						querySlice := util.GenerateSQLSelectQuerySlice(join.Relation, util.MergeSlices(true, join.Selects, constant.DEFAULT_JOIN_SELECT_COLUMNS))
						if query.Distinct {
							qj = qj.Distinct(querySlice)
						} else {
							qj = qj.Select(querySlice)
						}
					}
					if len(query.Searches) > 0 {
						queryString, bindValues := util.GenerateSQLWhereQueryStringAndBindValues(join.Relation, join.Searches)
						qj = qj.Where(queryString, bindValues...)
					}
					q = q.Joins(join.Relation, qj)
				} else {
					q = q.Joins(join.Relation)
				}
			}
		}
		if len(query.Orders) > 0 {
			queryString := util.GenerateSQLOrderQueryString(tableName, query.Orders)
			q = q.Order(queryString)
		}
		if len(query.Groups) > 0 {
			queryString := util.GenerateSQLGroupQueryString(tableName, query.Groups)
			q = q.Group(queryString)
		}
		if query.Limit != 0 {
			q = q.Limit(query.Limit)
		}
		if query.Offset != 0 {
			q = q.Offset(query.Offset)
		}
		if query.WithDeleted {
			q = q.Unscoped()
		}
	}

	q = q.Find(&employees)

	if err := q.Error; err != nil {
		return nil, count, err
	}

	count = int(q.RowsAffected)

	return employees, count, nil
}
