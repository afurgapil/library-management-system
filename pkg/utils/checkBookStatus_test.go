package utils

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v4"
)

var dbConnection *pgx.Conn
//TODO move main
func TestMain(m *testing.M) {
	var err error
	dbConnection, err = pgx.Connect(context.Background(), "postgres://postgres:test1234@localhost:5432/librarymanagementsystem_test")
	if err != nil {
		panic(err)
	}
	defer dbConnection.Close(context.Background())

	code := m.Run()

	os.Exit(code)
}

func TestCheckBookStatus(t *testing.T) {
	tests := []struct {
		name    string
		bookID  string
		want    bool
		wantErr bool
	}{
		{
			name:    "Book available",
			bookID:  "c9b80392-a545-475f-9854-7d65269e0fc3",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Book is not available",
			bookID:  "67952f2f-6d49-4760-a615-534ab9a5556b",
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckBookStatus(dbConnection, tt.bookID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckBookStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckBookStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
