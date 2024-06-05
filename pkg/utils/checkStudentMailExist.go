package utils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func CheckStudentMailExist(db *pgx.Conn, mail string) (bool, error) {
	var mailExists int

	query := `SELECT COUNT(*) FROM student WHERE student_mail = $1`
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
