package testutils

import (
	"context"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

func SetupTestDataCreateEmployee(db *pgx.Conn, email string) error {
	_, err := db.Exec(context.Background(), `INSERT INTO employee (employee_id, employee_mail, employee_username, employee_phone_number, position, employee_password) 
              VALUES ($1, $2, $3, $4, $5, $6)`, "0", email, "a", "b", "c", "d")
	return err
}

func CleanupTestDataCreateEmployee(db *pgx.Conn, employeeMail string) error {
	_, err := db.Exec(context.Background(), `DELETE FROM employee WHERE employee_mail = $1`, employeeMail)
	return err
}

func SetupTestDataAuthenticateEmployee(db *pgx.Conn, email, username, phoneNumber, position, password string) error {
	_, err := db.Exec(context.Background(), `INSERT INTO employee (employee_id, employee_mail, employee_username, employee_phone_number, position, employee_password) 
              VALUES ($1, $2, $3, $4, $5, $6)`, "0", email, username, phoneNumber, position, password)
	return err
}

func CleanupTestDataAuthenticateEmployee(db *pgx.Conn, employeeMail string) error {
	_, err := db.Exec(context.Background(), `DELETE FROM employee WHERE employee_mail = $1`, employeeMail)
	return err
}

func SetupTestDataDeleteEmployee(db *pgx.Conn, id string) error {
	_, err := db.Exec(context.Background(), `INSERT INTO employee (employee_id, employee_mail, employee_username, employee_phone_number, position, employee_password) 
              VALUES ($1, $2, $3, $4, $5, $6)`, id, "email", "a", "b", "c", "d")
	return err
}

func CleanupTestDataDeleteEmployee(db *pgx.Conn, employeeID string) error {
	_, err := db.Exec(context.Background(), `DELETE FROM employee WHERE employee_id = $1`, employeeID)
	return err
}
