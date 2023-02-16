package util

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/rodericusifo/echo-template/internal/pkg/config"
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/internal/pkg/request"
)

func GetPortApp() string {
	return strconv.Itoa(config.Env.ServerPort)
}

func GetHostApp() string {
	return config.Env.ServerHost
}

func GetMysqlDBConnection() config.MysqlDBSQLConnection {
	return config.MysqlDBSQL
}

func GetPostgresDBConnection() config.PostgresDBSQLConnection {
	return config.PostgresDBSQL
}

func GetRedisDBCacheConnection() config.RedisDBCacheConnection {
	return config.RedisDBCache
}

func GetJWTAuthConfig() config.JWTAuthConfig {
	return config.JWTAuth
}

func GetRequestUser(c echo.Context) *request.RequestUser {
	return c.Get(constant.C_KEY_REQUEST_USER).(*request.RequestUser)
}
