package employee

import (
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/internal/pkg/util"
)

var (
	employeeDatabaseSQLRepository IEmployeeDatabaseSQLRepository
	mockQuery                     sqlmock.Sqlmock
)

var (
	mockDate, mockBirthday time.Time
	mockUUID, mockAddress  string
	mockAge                int
)

func SetupTestPostgresEmployeeDatabaseSQLRepository() {
	dialect := constant.POSTGRES
	db, mock := util.MockConnectionDatabaseSQL(dialect)

	employeeDatabaseSQLRepository = InitPostgresEmployeeDatabaseSQLRepository(db)
	mockQuery = mock

	layoutFormat := "2006-01-02 15:04:05"

	valueDate := "2015-09-02 08:04:00"
	mockDate, _ = time.Parse(layoutFormat, valueDate)

	valueBirthday := "1999-03-12 00:00:00"
	mockBirthday, _ = time.Parse(layoutFormat, valueBirthday)

	mockUUID = "ac0d6ce3-ff02-4024-896b-ea0ceba32182"
	mockAddress = "Street A, City B"
	mockAge = 25
}
