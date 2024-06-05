package utils

import (
	"testing"

	"github.com/afurgapil/library-management-system/pkg/entities"
	testutils "github.com/afurgapil/library-management-system/pkg/testUtils"
)

func TestUpdateLimit(t *testing.T) {
	type args struct {
		studentID  string
		updateType string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid Borrow Update",
			args: args{
				studentID:  "daadfed9-c2b3-40a2-9c87-b0d8130adcc5",
				updateType: "borrow",
			},
			wantErr: false,
		},
		{
			name: "Valid Delivery Update",
			args: args{
				studentID:  "daadfed9-c2b3-40a2-9c87-b0d8130adcc5",
				updateType: "delivery",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			student := &entities.Student{
				StudentID:       tt.args.studentID,
				StudentMail:     "student_mail",
				StudentPassword: "student_password",
				Debit:           20,
				BookLimit:       20,
				IsBanned:        false,
			}
			if err := testutils.SetupTestDataStudent(dbConnection, student); err != nil {
				t.Fatalf("Failed to set up test data: %v", err)
			}

			defer testutils.CleanupTestDataStudent(dbConnection)

			if err := UpdateLimit(dbConnection, tt.args.studentID, tt.args.updateType); (err != nil) != tt.wantErr {
				t.Errorf("UpdateLimit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
