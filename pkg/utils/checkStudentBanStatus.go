package utils

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

func CheckStudentBanStatus(db *pgx.Conn, email string) (bool, error) {
	var studentStatus bool = false

	query := `SELECT is_banned FROM student WHERE student_mail = $1`
	err := db.QueryRow(context.Background(), query, email).Scan(&studentStatus)
	if err != nil {
		return false, fmt.Errorf("error fetching student ban status: %w", err)
	}

	return studentStatus, err
}
