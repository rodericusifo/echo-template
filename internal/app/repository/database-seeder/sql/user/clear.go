package user

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/pkg/types"
	"github.com/rodericusifo/echo-template/pkg/util"

	log "github.com/sirupsen/logrus"
)

func (s *UserDatabaseSeederSQLRepository) Clear(db *gorm.DB) error {
	users := make([]*sql.User, 0)
	tableName := sql.User{}.TableName()

	q := db.Table(tableName)

	querySlice := util.GenerateSQLSelectQuerySlice(
		tableName,
		[]types.SelectOperation{
			{Field: "id"},
		},
	)
	q = q.Select(querySlice)

	queryString, bindValues := util.GenerateSQLWhereQueryStringAndBindValues(
		tableName,
		[][]types.SearchOperation{
			{
				{Field: "role", Operator: "=", Value: constant.ADMIN},
			},
		},
	)
	q = q.Where(queryString, bindValues...)

	if err := q.Find(&users).Error; err != nil {
		return err
	}

	for _, user := range users {
		if err := db.Table(tableName).Delete(user).Error; err != nil {
			log.WithFields(log.Fields{
				"message": "delete user fail",
				"detail":  err,
			}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [CLEAR]")
			continue
		}
	}

	return nil
}
