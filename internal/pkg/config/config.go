package config

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/internal/pkg/migration"
	"github.com/rodericusifo/echo-template/internal/pkg/types"

	echojwt "github.com/labstack/echo-jwt/v4"
	log "github.com/sirupsen/logrus"
)

var (
	Env EnvConfig
)

var (
	MysqlDBSQL    MysqlDBSQLConnection
	PostgresDBSQL PostgresDBSQLConnection
)
var (
	RedisDBCache RedisDBCacheConnection
)

var (
	JWTAuth JWTAuthConfig
)

func ConfigureEnv() {
	var (
		environment = flag.String("env", "", "input the environment type")
	)

	flag.Parse()

	switch constant.EnvironmentType(*environment) {
	case constant.DEV:
		viper.SetConfigFile("./environments/dev.application.env")
	case constant.DOCKER:
		viper.SetConfigFile("./environments/docker.application.env")
	default:
		log.WithFields(log.Fields{
			"message": "set env fail",
			"detail":  errors.New("input environment type [ dev | docker ]"),
		}).Panic("[CONFIGURE ENV]")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.WithFields(log.Fields{
			"message": "read env fail",
			"detail":  err,
		}).Panic("[CONFIGURE ENV]")
	}

	var env EnvConfig
	if err := viper.Unmarshal(&env); err != nil {
		log.WithFields(log.Fields{
			"message": "load env fail",
			"detail":  err,
		}).Panic("[CONFIGURE ENV]")
	}
	log.WithFields(log.Fields{
		"message": "load env success",
	}).Infoln("[CONFIGURE ENV]")

	Env = env
}

func ConfigureDatabaseSQL(dialect constant.DialectDatabaseSQLType) {
	var (
		dbSQLConfig DBSQLConfig
		db          *gorm.DB
		err         error
	)

	switch dialect {
	case constant.POSTGRES:
		dbSQLConfig = DBSQLConfig{
			Host:              Env.DatabasePostgresHost,
			Port:              Env.DatabasePostgresPort,
			Name:              Env.DatabasePostgresName,
			Username:          Env.DatabasePostgresUsername,
			Password:          Env.DatabasePostgresPassword,
			TimeZone:          Env.DatabasePostgresTimeZone,
			ConnectionTimeout: Env.DatabasePostgresConnectionTimeout,
			MaxIdleConnection: Env.DatabasePostgresMaxIdleConnection,
			MaxOpenConnection: Env.DatabasePostgresMaxOpenConnection,
			DebugMode:         Env.DatabasePostgresDebugMode,
		}
	case constant.MYSQL:
		dbSQLConfig = DBSQLConfig{
			Host:              Env.DatabaseMysqlHost,
			Port:              Env.DatabaseMysqlPort,
			Name:              Env.DatabaseMysqlName,
			Username:          Env.DatabaseMysqlUsername,
			Password:          Env.DatabaseMysqlPassword,
			ConnectionTimeout: Env.DatabaseMysqlConnectionTimeout,
			MaxIdleConnection: Env.DatabaseMysqlMaxIdleConnection,
			MaxOpenConnection: Env.DatabaseMysqlMaxOpenConnection,
			DebugMode:         Env.DatabaseMysqlDebugMode,
		}
	}

	cfg := &gorm.Config{
		Logger: logger.Default,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	}

	if dbSQLConfig.DebugMode {
		cfg.Logger = logger.Default.LogMode(logger.Info)
	}

	switch dialect {
	case constant.POSTGRES:
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
			dbSQLConfig.Host,
			dbSQLConfig.Username,
			dbSQLConfig.Password,
			dbSQLConfig.Name,
			dbSQLConfig.Port,
			dbSQLConfig.TimeZone)
		db, err = gorm.Open(postgres.Open(dsn), cfg)
		if err != nil {
			log.WithFields(log.Fields{
				"message": fmt.Sprintf("connect to database sql %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE SQL]")
		}
		log.WithFields(log.Fields{
			"message": fmt.Sprintf("connect to database sql %s success", dialect),
		}).Infoln("[CONFIGURE DATABASE SQL]")

		// Auto Migration Models
		db.AutoMigrate(migration.AutoMigrateModelList...)

		sqlDb, err := db.DB()
		if err != nil {
			log.WithFields(log.Fields{
				"message": fmt.Sprintf("set up database sql %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE SQL]")
		}
		log.WithFields(log.Fields{
			"message": fmt.Sprintf("set up database sql %s success", dialect),
		}).Infoln("[CONFIGURE DATABASE SQL]")
		sqlDb.SetConnMaxIdleTime(dbSQLConfig.ConnectionTimeout)
		sqlDb.SetMaxIdleConns(dbSQLConfig.MaxIdleConnection)
		sqlDb.SetMaxOpenConns(dbSQLConfig.MaxOpenConnection)

		PostgresDBSQL = db
	case constant.MYSQL:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbSQLConfig.Username,
			dbSQLConfig.Password,
			dbSQLConfig.Host,
			dbSQLConfig.Port,
			dbSQLConfig.Name)
		db, err = gorm.Open(mysql.Open(dsn), cfg)
		if err != nil {
			log.WithFields(log.Fields{
				"message": fmt.Sprintf("connect to database sql %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE SQL]")
		}
		log.WithFields(log.Fields{
			"message": fmt.Sprintf("connect to database sql %s success", dialect),
		}).Infoln("[CONFIGURE DATABASE SQL]")

		// Auto Migration Models
		db.AutoMigrate(migration.AutoMigrateModelList...)

		sqlDb, err := db.DB()
		if err != nil {
			log.WithFields(log.Fields{
				"message": fmt.Sprintf("set up database sql %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE SQL]")
		}
		log.WithFields(log.Fields{
			"message": fmt.Sprintf("set up database sql %s success", dialect),
		}).Infoln("[CONFIGURE DATABASE SQL]")
		sqlDb.SetConnMaxIdleTime(dbSQLConfig.ConnectionTimeout)
		sqlDb.SetMaxIdleConns(dbSQLConfig.MaxIdleConnection)
		sqlDb.SetMaxOpenConns(dbSQLConfig.MaxOpenConnection)

		MysqlDBSQL = db
	}
}

func ConfigureDatabaseCache(dialect constant.DialectDatabaseCacheType) {
	switch dialect {
	case constant.REDIS:
		rdb := redis.NewClient(&redis.Options{
			Addr:     Env.DatabaseCacheRedisAddress,
			Password: Env.DatabaseCacheRedisPassword,
			DB:       Env.DatabaseCacheRedisDatabase,
		})
		ctx := context.Background()
		if err := rdb.Ping(ctx).Err(); err != nil {
			log.WithFields(log.Fields{
				"message": fmt.Sprintf("connect to database cache %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE CACHE]")
		}
		log.WithFields(log.Fields{
			"message": fmt.Sprintf("connect to database cache %s success", dialect),
		}).Infoln("[CONFIGURE DATABASE CACHE]")
		RedisDBCache = rdb
	}
}

func ConfigureAuth() {
	jwtConfig := &echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(types.JwtCustomClaims)
		},
		SigningKey: []byte(Env.JWTSecretKey),
	}
	JWTAuth = jwtConfig
}

func ConfigureLog() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
		TimestampFormat:        "2006-01-02 15:04:05 MST",
	})
	log.WithFields(log.Fields{
		"message": "setting log success",
	}).Infoln("[CONFIGURE LOG]")
}
