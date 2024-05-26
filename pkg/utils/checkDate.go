package utils

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
)

func CheckDate( db *pgx.Conn, borrowID string) (int, error) {
	var deliveryDate string
	query := `SELECT delivery_date FROM book_borrow WHERE borrow_id = $1`
	err := db.QueryRow(context.Background(), query, borrowID).Scan(&deliveryDate)
	if err != nil {
		return 0, err
	}

	layout := "2006-01-02"
	deliveryTime, err := time.Parse(layout, deliveryDate)
	if err != nil {
		return 0, err
	}

	currentTime := time.Now()

	duration := deliveryTime.Sub(currentTime).Hours() / 24
	daysRemaining := int(duration)

	return daysRemaining, nil
}
