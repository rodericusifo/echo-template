package user

import (
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/pkg/types"
	"github.com/rodericusifo/echo-template/pkg/util"
)

func (r *PostgresUserDatabaseSQLRepository) GetUser(query *types.Query) (*sql.User, error) {
	user := new(sql.User)
	tableName := sql.User{}.TableName()

	q := r.db.Table(tableName)

	if query != nil {
		if len(query.Selects) > 0 {
			querySlice := util.GenerateSQLSelectQuerySlice(tableName, util.MergeSlices(true, query.Selects, constant.DEFAULT_SELECT_COLUMNS))
			q = q.Select(querySlice)
		}
		if len(query.Searches) > 0 {
			queryString, bindValues := util.GenerateSQLWhereQueryStringAndBindValues(tableName, query.Searches)
			q = q.Where(queryString, bindValues...)
		}
	}

	if err := q.First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
