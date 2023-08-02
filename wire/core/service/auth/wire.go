//go:build wireinject
// +build wireinject

package auth

import (
	"github.com/google/wire"

	internal_pkg_util "github.com/rodericusifo/echo-template/internal/pkg/util"

	auth_service "github.com/rodericusifo/echo-template/internal/app/core/auth/service"

	user_resource "github.com/rodericusifo/echo-template/internal/app/core/user/resource"
	user_database_repository "github.com/rodericusifo/echo-template/internal/app/repository/database/sql/user"
)

func AuthService() auth_service.IAuthService {
	wire.Build(
		internal_pkg_util.GetPostgresDBConnection,
		user_database_repository.InitPostgresUserDatabaseSQLRepository,
		user_resource.InitUserResource,
		auth_service.InitAuthService,
	)
	return &auth_service.AuthService{}
}
