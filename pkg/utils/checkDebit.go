package utils

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func CheckDebit(db *pgx.Conn, studentID string) (bool,error)  {
	var studentDebit int
	query :=`SELECT debit FROM student WHERE student_id = $1`

	err := db.QueryRow(context.Background(),query,studentID).Scan(&studentDebit)
	if err != nil {
		return false,err
	}
	return studentDebit <= 0, nil
}