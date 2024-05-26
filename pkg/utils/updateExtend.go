package utils

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func UpdateExtend(db *pgx.Conn, borrowID string) (bool, error) {
	var extendStatus bool

	query := `UPDATE book_borrow SET is_extended = $1 WHERE borrow_id = $2 RETURNING is_extended`
	err := db.QueryRow(context.Background(), query, true, borrowID).Scan(&extendStatus)
	if err != nil {
		return false, fmt.Errorf("error updating is_extended for borrow %s: %w", borrowID, err)
	}

	return extendStatus, nil
}
