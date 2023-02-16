package user

import (
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"

	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/internal/pkg/util"

	gorm_seeder "github.com/kachit/gorm-seeder"
)

var (
	userDatabaseSeederSQLRepository *UserDatabaseSeederSQLRepository
	mockQuery                       sqlmock.Sqlmock
	mockDB                          *gorm.DB
)

var (
	mockDate         time.Time
	mockHashPassword string
)

func SetupTestUserDatabaseSeederSQLRepository() {
	dialect := constant.POSTGRES
	db, mock := util.MockConnectionDatabaseSQL(dialect)

	userDatabaseSeederSQLRepository = InitUserDatabaseSeederSQLRepository(gorm_seeder.SeederConfiguration{})
	mockQuery = mock
	mockDB = db

	layoutFormat := "2006-01-02 15:04:05"
	value := "2015-09-02 08:04:00"
	mockDate, _ = time.Parse(layoutFormat, value)

	mockHashPassword = "$2y$14$rnbG3JhbftD.iQV0QRf5GeNI/XlI85KF2kzrf4hnOs48cSoqPvsmG"
}
