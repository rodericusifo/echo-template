package resource

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
)

func init() {
	SetupTestEmployeeResource()
}

func TestEmployeeResource_CreateEmployee(t *testing.T) {
	type (
		args struct {
			payload *sql.Employee
		}
		result struct {
			err error
		}
	)

	testCases := []struct {
		desc   string
		input  args
		output result
		before func()
		after  func()
	}{
		{
			desc: "[ERROR]_because_something_error_happens",
			input: args{
				payload: &sql.Employee{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Address:  &mockAddress,
					Age:      &mockAge,
					Birthday: &mockBirthday,
					UserID:   1,
				},
			},
			output: result{
				err: errors.New("error something"),
			},
			before: func() {
				{
					var (
						arg1 *sql.Employee = &sql.Employee{
							Name:     "Ifo",
							Email:    "Ifo@gmail.com",
							Address:  &mockAddress,
							Age:      &mockAge,
							Birthday: &mockBirthday,
							UserID:   1,
						}
					)
					var (
						err error = errors.New("error something")
					)
					mockEmployeeDatabaseSQLRepository.On("CreateEmployee", arg1).Return(err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_create_employee",
			input: args{
				payload: &sql.Employee{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Address:  &mockAddress,
					Age:      &mockAge,
					Birthday: &mockBirthday,
					UserID:   1,
				},
			},
			output: result{
				err: nil,
			},
			before: func() {
				{
					var (
						arg1 *sql.Employee = &sql.Employee{
							Name:     "Ifo",
							Email:    "Ifo@gmail.com",
							Address:  &mockAddress,
							Age:      &mockAge,
							Birthday: &mockBirthday,
							UserID:   1,
						}
					)
					var (
						err error = nil
					)
					mockEmployeeDatabaseSQLRepository.On("CreateEmployee", arg1).Return(err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			err := employeeResource.CreateEmployee(tC.input.payload)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
