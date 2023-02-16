package employee

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/internal/pkg/config"
	"github.com/rodericusifo/echo-template/pkg/types"
)

type IEmployeeDatabaseSQLRepository interface {
	CreateEmployee(payload *sql.Employee) error
	UpdateEmployee(payload *sql.Employee) error
	DeleteEmployee(payload *sql.Employee) error
	GetListEmployeeAndCount(query *types.Query) ([]*sql.Employee, int, error)
	GetEmployee(query *types.Query) (*sql.Employee, error)
	CountAllEmployee(query *types.Query) (int, error)
}

type PostgresEmployeeDatabaseSQLRepository struct {
	db *gorm.DB
}

func InitPostgresEmployeeDatabaseSQLRepository(db config.PostgresDBSQLConnection) IEmployeeDatabaseSQLRepository {
	return &PostgresEmployeeDatabaseSQLRepository{
		db: db,
	}
}
