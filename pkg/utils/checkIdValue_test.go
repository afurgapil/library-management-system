package utils

import (
	"testing"

	"github.com/afurgapil/library-management-system/pkg/entities"
	testutils "github.com/afurgapil/library-management-system/pkg/testUtils"
)

func TestCheckIdValue(t *testing.T) {
	type args struct {
		tableName  string
		columnName string
		idValue    string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Student exist available",
			args: args{
				tableName:  "student",
				columnName: "student_id",
				idValue:    "daadfed9-c2b3-40a2-9c87-b0d8130adcc5",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Student is not exist",
			args: args{
				tableName:  "student",
				columnName: "student_id",
				idValue:    "realkingrobbstark",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.CleanupTestDataStudent(dbConnection)

			if tt.want {
				if err := testutils.SetupTestDataStudent(dbConnection, &entities.Student{
					StudentID:       tt.args.idValue,
					StudentMail:     "student_mail",
					StudentPassword: "student_password",
					Debit:           20,
					BookLimit:       20,
					IsBanned:        false,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			}

			got, err := CheckIdValue(dbConnection, tt.args.tableName, tt.args.columnName, tt.args.idValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckIdValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckIdValue() = %v, want %v", got, tt.want)
			}

			defer testutils.CleanupTestDataStudent(dbConnection)
		})
	}
}
