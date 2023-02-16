package util

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rodericusifo/echo-template/internal/pkg/constant"

	log "github.com/sirupsen/logrus"
)

func MockConnectionDatabaseSQL(dialect constant.DialectDatabaseSQLType) (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp),
	)
	if err != nil {
		log.WithFields(log.Fields{
			"message": fmt.Sprintf("connect to mock database sql %s failed", dialect),
			"detail":  err,
		}).Panic("[MOCK CONNECTION DATABASE SQL]")
	}

	var dialector gorm.Dialector
	switch dialect {
	case constant.POSTGRES:
		dialector = postgres.New(postgres.Config{
			Conn:       sqlDB,
			DriverName: string(constant.POSTGRES),
		})
	case constant.MYSQL:
		dialector = mysql.New(mysql.Config{
			Conn:                      sqlDB,
			DriverName:                string(constant.MYSQL),
			SkipInitializeWithVersion: true,
		})
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.WithFields(log.Fields{
			"message": fmt.Sprintf("connect to mock database sql %s failed", dialect),
			"detail":  err,
		}).Panic("[MOCK CONNECTION DATABASE SQL]")
	}

	return db, mock
}
