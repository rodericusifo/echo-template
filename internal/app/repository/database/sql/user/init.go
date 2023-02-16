package user

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/internal/pkg/config"
	"github.com/rodericusifo/echo-template/pkg/types"
)

type IUserDatabaseSQLRepository interface {
	CreateUser(payload *sql.User) error
	GetUser(query *types.Query) (*sql.User, error)
}

type PostgresUserDatabaseSQLRepository struct {
	db *gorm.DB
}

func InitPostgresUserDatabaseSQLRepository(db config.PostgresDBSQLConnection) IUserDatabaseSQLRepository {
	return &PostgresUserDatabaseSQLRepository{
		db: db,
	}
}
