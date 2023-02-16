package helper

import (
	"github.com/rodericusifo/echo-template/internal/app/repository/database-seeder/sql/user"
	"github.com/rodericusifo/echo-template/internal/pkg/config"
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/internal/pkg/util"
)

func RunDatabaseSeederSQL(dialect constant.DialectDatabaseSQLType) {
	switch dialect {
	case constant.POSTGRES:
		user.ExecutePostgresUserDatabaseSeederRepository(config.Env.DatabaseSeederPostgresUserIsRebuildData, util.GetPostgresDBConnection())
	case constant.MYSQL:
	}
}
