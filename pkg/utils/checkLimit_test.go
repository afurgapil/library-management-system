package utils

import (
	"testing"

	"github.com/afurgapil/library-management-system/pkg/entities"
	testutils "github.com/afurgapil/library-management-system/pkg/testUtils"
)

func TestCheckLimit(t *testing.T) {
	tests := []struct {
		name      string
		studentID string
		bookLimit int
		want      bool
		wantErr   bool
	}{
		{
			name:      "Limit okay",
			studentID: "f6df1bd5-a4dc-4390-a1bc-b92d73dffec8",
			bookLimit: 3,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "No Limit",
			studentID: "63a963fd-2753-4a26-9ef0-26777c472d35",
			bookLimit: 0,
			want:      false,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.CleanupTestDataStudent(dbConnection)

			if err := testutils.SetupTestDataStudent(dbConnection, &entities.Student{
				StudentID:       tt.studentID,
				StudentMail:     "student_mail",
				StudentPassword: "student_password",
				Debit:           20,
				BookLimit:       tt.bookLimit,
				IsBanned:        false,
			}); err != nil {
				t.Fatalf("Failed to set up test data: %v", err)
			}

			got, err := CheckLimit(dbConnection, tt.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckLimit() = %v, want %v", got, tt.want)
			}

			defer testutils.CleanupTestDataStudent(dbConnection)
		})
	}
}
