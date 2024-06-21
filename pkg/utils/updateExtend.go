package utils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

func UpdateExtend(db *pgx.Conn, borrowID string) (bool, error) {
	query, err := db.Exec(context.Background(), `
		UPDATE book_borrow
		SET is_extended = true
		WHERE borrow_id = $1
	`, borrowID)

	if err != nil {
		return false, err
	}

	if query.RowsAffected() == 0 {
		return false, errors.New("borrow ID not found")
	}

	return true, nil
}
