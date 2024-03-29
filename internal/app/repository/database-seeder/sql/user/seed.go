package user

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/internal/pkg/validator"
	"github.com/rodericusifo/echo-template/pkg/types"
	"github.com/rodericusifo/echo-template/pkg/util"

	mocks_pkg "github.com/rodericusifo/echo-template/mocks-pkg"
	log "github.com/sirupsen/logrus"
)

func (s *UserDatabaseSeederSQLRepository) Seed(db *gorm.DB) error {
	users := make([]*sql.User, 0)
	for _, UserSeed := range UserSeedData {
		err := validator.ValidatePayload(UserSeed)
		if err != nil {
			log.WithFields(log.Fields{
				"message": fmt.Sprintf("validation failed: user with xid %s", UserSeed.XID),
				"detail":  err,
			}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [SEED]")
			continue
		}

		user := new(sql.User)
		tableName := sql.User{}.TableName()

		q := db

		query := &types.Query{
			Selects: []types.SelectOperation{
				{Field: "id"},
			},
			Searches: [][]types.SearchOperation{
				{
					{Field: "xid", Operator: "=", Value: UserSeed.XID},
					{Field: "email", Operator: "=", Value: UserSeed.Email},
				},
			},
		}

		q = util.BuildQuery(tableName, q, query)

		err = q.Table(tableName).First(user).Error

		if err != nil && err != gorm.ErrRecordNotFound {
			log.WithFields(log.Fields{
				"message": "get user fail",
				"detail":  err,
			}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [SEED]")
			continue
		}
		if user.ID != 0 {
			log.WithFields(log.Fields{
				"message": fmt.Sprintf("user with xid %s and email %s already registered", UserSeed.XID, UserSeed.Email),
			}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [SEED]")
			continue
		}

		hashedPassword, err := mocks_pkg.GenerateHashPasswordUtil(UserSeed.Password)
		if err != nil {
			log.WithFields(log.Fields{
				"message": fmt.Sprintf("hash password fail: user with xid %s", UserSeed.XID),
				"detail":  err,
			}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [SEED]")
			continue
		}

		users = append(users, &sql.User{
			XID:      UserSeed.XID,
			Name:     UserSeed.Name,
			Email:    UserSeed.Email,
			Password: hashedPassword,
			Role:     UserSeed.Role,
		})
	}
	return db.CreateInBatches(users, len(users)).Error
}
