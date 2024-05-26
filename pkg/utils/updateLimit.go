package utils

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func UpdateLimit(db *pgx.Conn, studentID string, updateType string) (error) {
    var bookLimit int
    var newLimit int
    query := `SELECT book_limit from student WHERE student_id = $1`

    err := db.QueryRow(context.Background(), query, studentID).Scan(&bookLimit)
    if err != nil {
        return fmt.Errorf("error fetching book limit for student %s: %w", studentID, err)
    }

    switch updateType {
    case "borrow":
        newLimit = bookLimit - 1
    case "delivery":
        newLimit = bookLimit + 1
    default:
        return fmt.Errorf("invalid update type: %s", updateType)
    }

    updateQuery := `UPDATE student SET book_limit = $1 WHERE student_id = $2`
    _, err = db.Exec(context.Background(), updateQuery, newLimit, studentID)
    if err != nil {
        return fmt.Errorf("error updating book_limit for student %s: %w", studentID, err)
    }

    return nil
}
