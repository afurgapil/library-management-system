package utils

import (
	"testing"

	"github.com/afurgapil/library-management-system/pkg/entities"
	testutils "github.com/afurgapil/library-management-system/pkg/testUtils"
)

func TestCheckDebit(t *testing.T) {
	tests := []struct {
		name      string
		studentID string
		debit     int
		want      bool
		wantErr   bool
	}{
		{
			name:      "Debit = 0",
			studentID: "f6df1bd5-a4dc-4390-a1bc-b92d73dffec8",
			debit:     0,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "Debit > 0",
			studentID: "0d7d1132-c42c-4e64-9ff6-970310b14e28",
			debit:     10,
			want:      false,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			student := &entities.Student{
				StudentID:       tt.studentID,
				StudentMail:     "student_mail",
				StudentPassword: "student_password",
				Debit:           tt.debit,
				BookLimit:       20,
				IsBanned:        false,
			}
			if err := testutils.SetupTestDataStudent(dbConnection, student); err != nil {
				t.Fatalf("Failed to set up student data: %v", err)
			}
			defer testutils.CleanupTestDataStudent(dbConnection)
			got, err := CheckDebit(dbConnection, tt.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckDebit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckDebit() = %v, want %v", got, tt.want)
			}
		})
	}
}
