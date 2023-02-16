package config

import (
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type (
	MysqlDBSQLConnection    *gorm.DB
	PostgresDBSQLConnection *gorm.DB
)

type (
	RedisDBCacheConnection *redis.Client
)

type DBSQLConfig struct {
	Host              string
	Port              string
	Name              string
	Username          string
	Password          string
	TimeZone          string
	ConnectionTimeout time.Duration
	MaxIdleConnection int
	MaxOpenConnection int
	DebugMode         bool
}
