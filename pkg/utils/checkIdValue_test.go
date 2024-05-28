package utils

import (
	"testing"
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
				idValue:    "c96d572b-1f54-438a-8a00-6191f09645d9",
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
			got, err := CheckIdValue(dbConnection, tt.args.tableName, tt.args.columnName, tt.args.idValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckIdValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckIdValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
