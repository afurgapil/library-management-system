package utils

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func UpdateDebit(db *pgx.Conn, studentID string, additionalDebit int) error {
	var currentDebit int
	query := `SELECT debit FROM student WHERE student_id = $1`
	err := db.QueryRow(context.Background(), query, studentID).Scan(&currentDebit)
	if err != nil {
		return fmt.Errorf("error fetching current debit for student %s: %w", studentID, err)
	}

	newDebit := currentDebit + additionalDebit

	updateQuery := `UPDATE student SET debit = $1 WHERE student_id = $2`
	_, err = db.Exec(context.Background(), updateQuery, newDebit, studentID)
	if err != nil {
		return fmt.Errorf("error updating debit for student %s: %w", studentID, err)
	}

	return nil
}
