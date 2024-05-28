package utils

import (
	"testing"
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
				studentID      :"f6df1bd5-a4dc-4390-a1bc-b92d73dffec8",
				additionalDebit :1,
			},
			wantErr: false,
		},
			{
			name: "Update -1",
			args: args{
				studentID      :"f6df1bd5-a4dc-4390-a1bc-b92d73dffec8",
				additionalDebit :-1,
			},
			wantErr: false,
		},
		{
			name: "Update 1 && Invalid User ID",
			args: args{
				studentID      :"invalid-user-id",
				additionalDebit :1,
			},
			wantErr: true,
		},
		{
			name: "Update -1 && Invalid User ID",
			args: args{
				studentID      :"invalid-user-id",
				additionalDebit :-1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateDebit(dbConnection, tt.args.studentID, tt.args.additionalDebit); (err != nil) != tt.wantErr {
				t.Errorf("UpdateDebit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
