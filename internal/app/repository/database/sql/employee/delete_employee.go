package employee

import (
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
)

func (r *PostgresEmployeeDatabaseSQLRepository) DeleteEmployee(payload *sql.Employee) error {
	employee := new(sql.Employee)
	tableName := sql.Employee{}.TableName()

	q := r.db.Table(tableName)

	if payload != nil {
		employee = payload
	}

	if err := q.Delete(employee).Error; err != nil {
		return err
	}

	return nil
}
