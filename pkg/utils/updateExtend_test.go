package utils

import (
	"testing"

	testutils "github.com/afurgapil/library-management-system/pkg/testUtils"
)

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
			borrowID: "borrow_id",
			setup:    true,
			want:     true,
			wantErr:  false,
		}, {
			name:     "Invalid Borrow ID",
			borrowID: "borrow_id",
			setup:    false,
			want:     false,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer testutils.CleanupTestDataBook(dbConnection)
			defer testutils.CleanupTestDataStudent(dbConnection)
			defer testutils.CleanupTestDataBookBorrow(dbConnection)

			if tt.setup {
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

			got, err := UpdateExtend(dbConnection, tt.borrowID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateExtend() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateExtend() = %v, want %v", got, tt.want)
			}
			defer testutils.CleanupTestDataBook(dbConnection)
			defer testutils.CleanupTestDataStudent(dbConnection)
			defer testutils.CleanupTestDataBookBorrow(dbConnection)
		})
	}
}
