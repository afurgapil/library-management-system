package utils

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func CheckLimit(db *pgx.Conn, studentID string) (bool, error) {
    var studentBookLimit int
    query := `SELECT book_limit FROM student WHERE student_id = $1`
    
    err := db.QueryRow(context.Background(), query, studentID).Scan(&studentBookLimit)
    if err != nil {
        return false, err
    }
    return studentBookLimit > 0, nil
}
