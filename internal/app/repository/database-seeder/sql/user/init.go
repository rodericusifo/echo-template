package user

import (
	gorm_seeder "github.com/kachit/gorm-seeder"
)

type UserDatabaseSeederSQLRepository struct {
	gorm_seeder.SeederAbstract
}

func InitUserDatabaseSeederSQLRepository(cfg gorm_seeder.SeederConfiguration) *UserDatabaseSeederSQLRepository {
	return &UserDatabaseSeederSQLRepository{gorm_seeder.NewSeederAbstract(cfg)}
}
