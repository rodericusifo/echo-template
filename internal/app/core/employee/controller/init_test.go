package controller

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/rodericusifo/echo-template/internal/pkg/request"
	"github.com/rodericusifo/echo-template/mocks"
)

var (
	mockApp             *echo.Echo
	mockEmployeeService *mocks.IEmployeeService
	employeeController  *EmployeeController
)

var (
	mockBirthday, mockDate             time.Time
	mockAddress, mockUUID, mockUserXID string
	mockAge, mockAgeMinus              int
	mockReqUser                        *request.RequestUser
)

func SetupTestEmployeeController() {
	mockApp = echo.New()

	mockEmployeeService = new(mocks.IEmployeeService)

	employeeController = InitEmployeeController(mockEmployeeService)
	employeeController.Mount(mockApp.Group("/employees"))

	mockAddress = "20196 Morton Drive"
	mockAge = 24
	mockAgeMinus = -1

	layoutFormat := "2006-01-02 15:04:05"

	valueBirthday := "1999-08-02 08:04:00"
	mockBirthday, _ = time.Parse(layoutFormat, valueBirthday)

	valueDate := "2015-09-02 08:04:00"
	mockDate, _ = time.Parse(layoutFormat, valueDate)

	mockUUID = "ac0d6ce3-ff02-4024-896b-ea0ceba32182"
	mockReqUser = &request.RequestUser{
		ID: 1,
	}
}
