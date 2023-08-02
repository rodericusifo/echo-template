package user

import (
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/pkg/types"
	"github.com/rodericusifo/echo-template/pkg/util"
)

func (r *PostgresUserDatabaseSQLRepository) GetUser(query *types.Query) (*sql.User, error) {
	user := new(sql.User)
	tableName := sql.User{}.TableName()

	q := r.db

	if query != nil {
		q = util.BuildQuery(tableName, q, query)
	}

	if err := q.Table(tableName).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
