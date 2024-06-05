package utils

import (
	"testing"

	"github.com/afurgapil/library-management-system/pkg/entities"
	testutils "github.com/afurgapil/library-management-system/pkg/testUtils"
)

func TestUpdateDebit(t *testing.T) {
	type args struct {
		studentID       string
		additionalDebit int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Update 1",
			args: args{
				studentID:       "student_id",
				additionalDebit: 1,
			},
			wantErr: false,
		},
		{
			name: "Update -1",
			args: args{
				studentID:       "student_id",
				additionalDebit: -1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.CleanupTestDataStudent(dbConnection)
			if err := testutils.SetupTestDataStudent(dbConnection, &entities.Student{
				StudentID:       tt.args.studentID,
				StudentMail:     "student_mail",
				StudentPassword: "student_password",
				Debit:           20,
				BookLimit:       10,
				IsBanned:        false,
			}); err != nil {
				t.Fatalf("Failed to set up test data: %v", err)
			}
			if err := UpdateDebit(dbConnection, tt.args.studentID, tt.args.additionalDebit); (err != nil) != tt.wantErr {
				t.Errorf("UpdateDebit() error = %v, wantErr %v", err, tt.wantErr)
			}
			testutils.CleanupTestDataStudent(dbConnection)

		})
	}
}
