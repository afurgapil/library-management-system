package utils

import (
	"testing"

	"github.com/afurgapil/library-management-system/pkg/entities"
	testutils "github.com/afurgapil/library-management-system/pkg/testUtils"
)

func TestCheckStudentBanStatus(t *testing.T) {

	tests := []struct {
		name    string
		email   string
		want    bool
		wantErr bool
	}{
		{
			name:    "Not Banned",
			email:   "unlimitedLimit@gmail.com",
			want:    false,
			wantErr: false,
		}, {
			name:    "Banned",
			email:   "banned@gmail.com",
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.CleanupTestDataStudent(dbConnection)
			if err := testutils.SetupTestDataStudent(dbConnection, &entities.Student{
				StudentID:       "student_id",
				StudentMail:     tt.email,
				StudentPassword: "student_password",
				Debit:           20,
				BookLimit:       10,
				IsBanned:        tt.want,
			}); err != nil {
				t.Fatalf("Failed to set up test data: %v", err)
			}
			got, err := CheckStudentBanStatus(dbConnection, tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckStudentBanStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckStudentBanStatus() = %v, want %v", got, tt.want)
			}
			defer testutils.CleanupTestDataStudent(dbConnection)

		})
	}
}
