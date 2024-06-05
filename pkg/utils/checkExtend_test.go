package utils

import (
	"testing"

	"github.com/afurgapil/library-management-system/pkg/entities"
	testutils "github.com/afurgapil/library-management-system/pkg/testUtils"
)

func TestCheckExtend(t *testing.T) {
	tests := []struct {
		name           string
		borrowID       string
		isExtended     bool
		want           bool
		wantErr        bool
		wantErrMessage string
	}{
		{
			name:           "Extend available",
			borrowID:       "borrow_book_1",
			isExtended:     false,
			want:           false,
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name:           "Extend is not available",
			borrowID:       "borrow_book_2",
			isExtended:     true,
			want:           true,
			wantErr:        true,
			wantErrMessage: "date extended already",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.CleanupTestDataBook(dbConnection)
			testutils.CleanupTestDataStudent(dbConnection)
			testutils.CleanupTestDataBookBorrow(dbConnection)

			borrowedBook := &entities.BorrowedBook{
				BorrowID:     tt.borrowID,
				StudentID:    "student_id",
				BookID:       "book_id",
				BorrowDate:   "2024-01-01",
				DeliveryDate: "2024-01-10",
				IsExtended:   tt.isExtended,
			}
			if err := testutils.SetupTestDataBook(dbConnection, testutils.ExampleBook); err != nil {
				t.Fatalf("Failed to set up test data: %v", err)
			}
			if err := testutils.SetupTestDataStudent(dbConnection, testutils.ExampleStudent); err != nil {
				t.Fatalf("Failed to set up test data: %v", err)
			}
			if err := testutils.SetupTestDataBookBorrow(dbConnection, borrowedBook); err != nil {
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

			defer testutils.CleanupTestDataBook(dbConnection)
			defer testutils.CleanupTestDataStudent(dbConnection)
			defer testutils.CleanupTestDataBookBorrow(dbConnection)
		})
	}
}
