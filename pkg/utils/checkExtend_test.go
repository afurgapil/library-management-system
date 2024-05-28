package utils

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4"
)

func setupTestDataCheckExtend(db *pgx.Conn,borrowID string, isExtended bool) error  {
	_, err := db.Exec(context.Background(),`UPDATE book_borrow SET is_extended =$1 WHERE borrow_id = $2`,isExtended,borrowID)
	
	return err
}

func TestCheckExtend(t *testing.T) {
	tests := []struct {
		name      string
		borrowID  string
		want      bool
		wantErr   bool
		wantErrMessage string
	}{
		{
			name:      "Extend available",
			borrowID:  "4f0e40ba-e600-421e-baa1-19fa9595fe54",
			want:      false,
			wantErr:   false,
			wantErrMessage: "",
		},
		{
			name:      "Extend is not available",
			borrowID:  "046eec7f-67a1-45ec-9bed-29b88ea60400",
			want:      true,
			wantErr:   true,
			wantErrMessage: "date extended already",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err :=setupTestDataCheckExtend(dbConnection,tt.borrowID,tt.want); err !=nil {
				t.Fatalf("Failed to set up test data: %v", err)
			}
			got, err := CheckExtend(dbConnection, tt.borrowID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckExtend() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err.Error() != tt.wantErrMessage {
				t.Errorf("CheckExtend() error message = %v, wantErrMessage %v", err.Error(), tt.wantErrMessage)
				return
			}
			if got != tt.want {
				t.Errorf("CheckExtend() = %v, want %v", got, tt.want)
			}
		})
	}
}
