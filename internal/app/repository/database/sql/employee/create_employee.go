package employee

import (
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
)

func (r *PostgresEmployeeDatabaseSQLRepository) CreateEmployee(payload *sql.Employee) error {
	employee := new(sql.Employee)
	tableName := sql.Employee{}.TableName()

	q := r.db

	if payload != nil {
		employee = payload
	}

	if err := q.Table(tableName).Save(employee).Error; err != nil {
		return err
	}

	return nil
}
