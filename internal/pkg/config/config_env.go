package config

import (
	"time"
)

type (
	IsRebuildDataDBSeederMysqlUser bool
)

type (
	IsRebuildDataDBSeederPostgresUser bool
)

type EnvConfig struct {
	// APPS
	AppsName    string `mapstructure:"APPS_NAME"`
	AppsVersion string `mapstructure:"APPS_VERSION"`

	// SERVER
	ServerHost string `mapstructure:"SERVER_HOST"`
	ServerPort int    `mapstructure:"SERVER_PORT"`

	// DATABASE SQL
	DatabaseMysqlHost              string        `mapstructure:"DATABASE_MYSQL_HOST"`
	DatabaseMysqlPort              string        `mapstructure:"DATABASE_MYSQL_PORT"`
	DatabaseMysqlName              string        `mapstructure:"DATABASE_MYSQL_NAME"`
	DatabaseMysqlUsername          string        `mapstructure:"DATABASE_MYSQL_USERNAME"`
	DatabaseMysqlPassword          string        `mapstructure:"DATABASE_MYSQL_PASSWORD"`
	DatabaseMysqlConnectionTimeout time.Duration `mapstructure:"DATABASE_MYSQL_CONNECTION_TIMEOUT"`
	DatabaseMysqlMaxIdleConnection int           `mapstructure:"DATABASE_MYSQL_MAX_IDLE_CONNECTION"`
	DatabaseMysqlMaxOpenConnection int           `mapstructure:"DATABASE_MYSQL_MAX_OPEN_CONNECTION"`
	DatabaseMysqlDebugMode         bool          `mapstructure:"DATABASE_MYSQL_DEBUG_MODE"`

	DatabasePostgresHost              string        `mapstructure:"DATABASE_POSTGRES_HOST"`
	DatabasePostgresPort              string        `mapstructure:"DATABASE_POSTGRES_PORT"`
	DatabasePostgresName              string        `mapstructure:"DATABASE_POSTGRES_NAME"`
	DatabasePostgresUsername          string        `mapstructure:"DATABASE_POSTGRES_USERNAME"`
	DatabasePostgresPassword          string        `mapstructure:"DATABASE_POSTGRES_PASSWORD"`
	DatabasePostgresTimeZone          string        `mapstructure:"DATABASE_POSTGRES_TIME_ZONE"`
	DatabasePostgresConnectionTimeout time.Duration `mapstructure:"DATABASE_POSTGRES_CONNECTION_TIMEOUT"`
	DatabasePostgresMaxIdleConnection int           `mapstructure:"DATABASE_POSTGRES_MAX_IDLE_CONNECTION"`
	DatabasePostgresMaxOpenConnection int           `mapstructure:"DATABASE_POSTGRES_MAX_OPEN_CONNECTION"`
	DatabasePostgresDebugMode         bool          `mapstructure:"DATABASE_POSTGRES_DEBUG_MODE"`

	// DATABASE SEEDER SQL
	DatabaseSeederMysqlUserIsRebuildData IsRebuildDataDBSeederMysqlUser `mapstructure:"DATABASE_SEEDER_MYSQL_USER_IS_REBUILD_DATA"`

	DatabaseSeederPostgresUserIsRebuildData IsRebuildDataDBSeederPostgresUser `mapstructure:"DATABASE_SEEDER_POSTGRES_USER_IS_REBUILD_DATA"`

	// DATABASE CACHE
	DatabaseCacheRedisAddress  string `mapstructure:"DATABASE_CACHE_REDIS_ADDRESS"`
	DatabaseCacheRedisPassword string `mapstructure:"DATABASE_CACHE_REDIS_PASSWORD"`
	DatabaseCacheRedisDatabase int    `mapstructure:"DATABASE_CACHE_REDIS_DATABASE"`

	// PASSWORD HASHING
	PasswordHashingHashSalt int `mapstructure:"PASSWORD_HASHING_HASH_SALT"`

	// JWT
	JWTSecretKey       string        `mapstructure:"JWT_SECRET_KEY"`
	JWTExpiredDuration time.Duration `mapstructure:"JWT_EXPIRED_DURATION"`
}
