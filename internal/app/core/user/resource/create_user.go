package resource

import (
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
)

func (r *UserResource) CreateUser(payload *sql.User) error {
	return r.UserDatabaseSQLRepository.CreateUser(payload)
}
