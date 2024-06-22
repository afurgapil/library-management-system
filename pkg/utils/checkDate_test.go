package utils

import (
	"testing"
	"time"

	"github.com/afurgapil/library-management-system/pkg/entities"
	testutils "github.com/afurgapil/library-management-system/pkg/testUtils"
)

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
			borrowID:     "borrow_id_1",
			deliveryDate: time.Now().AddDate(0, 0, 30).Format("2006-01-02"),
			want:         29,
			wantErr:      false,
		},
		{
			name:         "0 days remaining",
			borrowID:     "borrow_id_2",
			deliveryDate: time.Now().Format("2006-01-02"),
			want:         0,
			wantErr:      false,
		},
		{
			name:         "<0 days remaining",
			borrowID:     "borrow_id_3",
			deliveryDate: time.Now().AddDate(0, 0, -30).Format("2006-01-02"),
			want:         -30,
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testutils.SetupTestDataBook(dbConnection, testutils.ExampleBook); err != nil {
				t.Fatalf("Failed to set up test data: %v", err)
			}
			if err := testutils.SetupTestDataStudent(dbConnection, testutils.ExampleStudent); err != nil {
				t.Fatalf("Failed to set up test data: %v", err)
			}

			borrowedBook := &entities.BorrowedBook{
				BorrowID:     tt.borrowID,
				StudentID:    "student_id",
				BookID:       "book_id",
				BorrowDate:   time.Now().AddDate(0, 0, -10).Format("2006-01-02"), // Örneğin 10 gün önce ödünç alındı
				DeliveryDate: tt.deliveryDate,
				IsExtended:   false,
			}

			if err := testutils.SetupTestDataBookBorrow(dbConnection, borrowedBook); err != nil {
				t.Fatalf("Failed to set up test data: %v", err)
			}

			defer testutils.CleanupTestDataBook(dbConnection)
			defer testutils.CleanupTestDataStudent(dbConnection)
			defer testutils.CleanupTestDataBookBorrow(dbConnection)

			got, err := CheckDate(dbConnection, tt.borrowID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckDate() error = %v, wantErr %v, testname: %v", err, tt.wantErr, tt.name)
				return
			}
			if got != tt.want {
				t.Errorf("CheckDate() = %v, want %v, testname %v", got, tt.want, tt.name)
			}
		})
	}
}
