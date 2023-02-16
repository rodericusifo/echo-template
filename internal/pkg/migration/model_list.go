package migration

import (
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
)

var (
	AutoMigrateModelList = []any{
		&sql.User{},
		&sql.Employee{},
	}
)
