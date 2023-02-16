package resource

import (
	"time"

	"github.com/rodericusifo/echo-template/mocks"
)

var (
	mockEmployeeDatabaseSQLRepository *mocks.IEmployeeDatabaseSQLRepository
	employeeResource                  IEmployeeResource
)

var (
	mockDate, mockBirthday time.Time
	mockUUID, mockAddress  string
	mockAge                int
)

func SetupTestEmployeeResource() {
	mockEmployeeDatabaseSQLRepository = new(mocks.IEmployeeDatabaseSQLRepository)

	employeeResource = InitEmployeeResource(mockEmployeeDatabaseSQLRepository)

	layoutFormat := "2006-01-02 15:04:05"
	value := "2015-09-02 08:04:00"
	mockDate, _ = time.Parse(layoutFormat, value)

	mockUUID = "ac0d6ce3-ff02-4024-896b-ea0ceba32182"
	// mockHashPassword = "$2y$14$rnbG3JhbftD.iQV0QRf5GeNI/XlI85KF2kzrf4hnOs48cSoqPvsmG"
}
