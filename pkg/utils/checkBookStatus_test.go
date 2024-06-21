package utils

import (
	"context"
	"os"
	"testing"

	testutils "github.com/afurgapil/library-management-system/pkg/testUtils"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

// Utils - DB Connection
var dbConnection *pgx.Conn

func TestMain(m *testing.M) {
	var err error
	dbConnection, err = testutils.InitializeDatabase()
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
			bookID:  "book_id",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Book is not available",
			bookID:  "book_id",
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "Book available":
				if err := testutils.SetupTestDataBook(dbConnection, testutils.ExampleBook); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataStudent(dbConnection, testutils.ExampleStudent); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			case "Book is not available":
				if err := testutils.SetupTestDataBook(dbConnection, testutils.ExampleBook); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataStudent(dbConnection, testutils.ExampleStudent); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataBookBorrow(dbConnection, testutils.ExampleBorrowedBook); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			}

			defer testutils.CleanupTestDataBook(dbConnection)
			defer testutils.CleanupTestDataStudent(dbConnection)
			defer testutils.CleanupTestDataBookBorrow(dbConnection)

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
