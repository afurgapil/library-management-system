package utils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

func CheckExtend(db *pgx.Conn, borrowID string) (bool, error) {
	var extendStatus bool

	query := `SELECT is_extended FROM book_borrow WHERE borrow_id = $1`
	err := db.QueryRow(context.Background(), query, borrowID).Scan(&extendStatus)
	if err != nil {
		return false, err
	}
	if extendStatus {
		return true, errors.New("date extended already")
	}
	return extendStatus, nil
}
