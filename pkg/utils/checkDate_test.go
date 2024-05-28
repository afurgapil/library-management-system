package utils

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
)

func setupTestDataCheckDate(db *pgx.Conn, borrowID string, deliveryDate string) error {
	_, err := db.Exec(context.Background(), `
		INSERT INTO book_borrow (borrow_id, student_id, book_id, borrow_date, delivery_date, is_extended) 
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (borrow_id) DO UPDATE SET delivery_date = EXCLUDED.delivery_date
	`, borrowID, "00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000000", "2024-05-27", deliveryDate, false)
	return err
}

func cleanupTestDataCheckDate(db *pgx.Conn, borrowID string) error {
	_, err := db.Exec(context.Background(), `
		DELETE FROM book_borrow WHERE borrow_id = $1
	`, borrowID)
	return err
}

func TestCheckDate(t *testing.T) {
	tests := []struct {
		name         string
		borrowID     string
		deliveryDate string
		want         int
		wantErr      bool
	}{
		{
			name:         ">1 days remaining",
			borrowID:     "d99bc27c-8ce7-42e7-9b33-46917189a01d",
			deliveryDate: time.Now().AddDate(0, 0, 30).Format("2006-01-02"),
			want:         29,
			wantErr:      false,
		},
		{
			name:         "0 days remaining",
			borrowID:     "83131602-6287-4ef7-ad1b-2749db97a45e",
			deliveryDate: time.Now().Format("2006-01-02"),
			want:         0,
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := setupTestDataCheckDate(dbConnection, tt.borrowID, tt.deliveryDate); err != nil {
				t.Fatalf("Failed to set up test data: %v", err)
			}
			defer cleanupTestDataCheckDate(dbConnection, tt.borrowID) 

			got, err := CheckDate(dbConnection, tt.borrowID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
