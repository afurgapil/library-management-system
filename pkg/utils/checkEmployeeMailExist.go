package utils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

func CheckEmployeeMailExist(db *pgx.Conn, mail string) (bool, error) {
	var mailExists int

	query := `SELECT COUNT(*) FROM employee WHERE employee_mail = $1`
	err := db.QueryRow(context.Background(), query, mail).Scan(&mailExists)
	if err != nil {
		return false, err
	}
	if mailExists == 0 {
		return false, nil
	} else {
		return true, errors.New("this mail has been used already")
	}

}
