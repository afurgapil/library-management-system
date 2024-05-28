package utils

import (
	"testing"
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
				studentID: "c96d572b-1f54-438a-8a00-6191f09645d9",
				updateType: "borrow",
			},
			wantErr: false,
		},	{
			name: "Invalid Borrow Update",
			args: args{
				studentID: "invalid-userID",
				updateType: "borrow",
			},
			wantErr: true,
		},	{
			name: "Valid Delivery Update",
			args: args{
				studentID: "c96d572b-1f54-438a-8a00-6191f09645d9",
				updateType: "delivery",
			},
			wantErr: false,
		},	{
			name: "Invalid Delivery Update",
			args: args{
				studentID: "invalid-userID",
				updateType: "delivery",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateLimit(dbConnection, tt.args.studentID, tt.args.updateType); (err != nil) != tt.wantErr {
				t.Errorf("UpdateLimit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
