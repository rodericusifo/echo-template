package resource

import (
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/internal/app/repository/database/sql/user"
	"github.com/rodericusifo/echo-template/pkg/types"
)

type IUserResource interface {
	CreateUser(payload *sql.User) error
	GetUser(query *types.Query) (*sql.User, error)
}

type UserResource struct {
	UserDatabaseSQLRepository user.IUserDatabaseSQLRepository
}

func InitUserResource(userDatabaseSQLRepository user.IUserDatabaseSQLRepository) IUserResource {
	return &UserResource{
		UserDatabaseSQLRepository: userDatabaseSQLRepository,
	}
}
