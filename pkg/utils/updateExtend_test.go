package utils

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4"
)

func setupTestDataUpdateExtend(db *pgx.Conn, borrowID string) error {
	_, err := db.Exec(context.Background(), `
		INSERT INTO book_borrow (borrow_id, student_id, book_id, borrow_date, delivery_date, is_extended) 
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (borrow_id) DO UPDATE SET delivery_date = EXCLUDED.delivery_date
	`, borrowID, "00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000000", "2024-05-27", "2024-06-27", false)
	return err
}

func cleanupTestDataUpdateExtend(db *pgx.Conn, borrowID string) error {
	_, err := db.Exec(context.Background(), `
		DELETE FROM book_borrow WHERE borrow_id = $1
	`, borrowID)
	return err
}

func TestUpdateExtend(t *testing.T) {
	tests := []struct {
		name     string
		borrowID string
		setup    bool
		want     bool
		wantErr  bool
	}{
		{
			name:     "Valid Borrow ID",
			borrowID: "d99bc27c-8ce7-42e7-9b33-46917189a01d",
			setup:    true,
			want:     true,
			wantErr:  false,
		}, {
			name:     "Invalid Borrow ID",
			borrowID: "invalid",
			setup:    false,
			want:     false,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup {
				if err := setupTestDataUpdateExtend(dbConnection, tt.borrowID); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				defer cleanupTestDataUpdateExtend(dbConnection, tt.borrowID)
			}

			got, err := UpdateExtend(dbConnection, tt.borrowID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateExtend() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateExtend() = %v, want %v", got, tt.want)
			}
		})
	}
}
