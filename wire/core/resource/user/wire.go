//go:build wireinject
// +build wireinject

package user

import (
	"github.com/google/wire"

	internal_pkg_util "github.com/rodericusifo/echo-template/internal/pkg/util"

	user_resource "github.com/rodericusifo/echo-template/internal/app/core/user/resource"
	user_database_repository "github.com/rodericusifo/echo-template/internal/app/repository/database/sql/user"
)

func UserResource() user_resource.IUserResource {
	wire.Build(
		internal_pkg_util.GetPostgresDBConnection,
		user_database_repository.InitPostgresUserDatabaseSQLRepository,
		user_resource.InitUserResource,
	)
	return &user_resource.UserResource{}
}
