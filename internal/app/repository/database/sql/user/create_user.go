package user

import (
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
)

func (r *PostgresUserDatabaseSQLRepository) CreateUser(payload *sql.User) error {
	user := new(sql.User)
	tableName := sql.User{}.TableName()

	q := r.db.Table(tableName)

	if payload != nil {
		user = payload
	}

	if err := q.Create(user).Error; err != nil {
		return err
	}

	return nil
}
