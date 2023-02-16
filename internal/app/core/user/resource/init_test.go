package resource

import (
	"time"

	"github.com/rodericusifo/echo-template/mocks"
)

var (
	mockUserDatabaseSQLRepository *mocks.IUserDatabaseSQLRepository
	userResource                  IUserResource
)

var (
	mockDate                   time.Time
	mockUUID, mockHashPassword string
)

func SetupTestUserResource() {
	mockUserDatabaseSQLRepository = new(mocks.IUserDatabaseSQLRepository)

	userResource = InitUserResource(mockUserDatabaseSQLRepository)

	layoutFormat := "2006-01-02 15:04:05"
	value := "2015-09-02 08:04:00"
	mockDate, _ = time.Parse(layoutFormat, value)

	mockUUID = "ac0d6ce3-ff02-4024-896b-ea0ceba32182"
	mockHashPassword = "$2y$14$rnbG3JhbftD.iQV0QRf5GeNI/XlI85KF2kzrf4hnOs48cSoqPvsmG"
}
