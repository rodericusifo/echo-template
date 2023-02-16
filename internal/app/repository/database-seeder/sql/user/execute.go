package user

import (
	"github.com/rodericusifo/echo-template/internal/pkg/config"

	gorm_seeder "github.com/kachit/gorm-seeder"
	log "github.com/sirupsen/logrus"
)

func ExecutePostgresUserDatabaseSeederRepository(isRebuildData config.IsRebuildDataDBSeederPostgresUser, db config.PostgresDBSQLConnection) {
	userDatabaseSeederSQLRepository := InitUserDatabaseSeederSQLRepository(gorm_seeder.SeederConfiguration{})
	seedersStack := gorm_seeder.NewSeedersStack(db)
	seedersStack.AddSeeder(userDatabaseSeederSQLRepository)

	if isRebuildData {
		err := seedersStack.Clear()
		if err != nil {
			log.WithFields(log.Fields{
				"message": "clear user fail",
				"detail":  err,
			}).Errorln("[EXECUTE POSTGRES USER DATABASE SEEDER REPOSITORY]")
			return
		}
		log.WithFields(log.Fields{
			"message": "clear user success",
		}).Infoln("[EXECUTE POSTGRES USER DATABASE SEEDER REPOSITORY]")
	}

	err := seedersStack.Seed()
	if err != nil {
		log.WithFields(log.Fields{
			"message": "seed user fail",
			"detail":  err,
		}).Errorln("[EXECUTE POSTGRES USER DATABASE SEEDER REPOSITORY]")
		return
	}
	log.WithFields(log.Fields{
		"message": "seed user success",
	}).Infoln("[EXECUTE POSTGRES USER DATABASE SEEDER REPOSITORY]")
}
