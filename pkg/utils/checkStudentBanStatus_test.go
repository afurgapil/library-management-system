package utils

import (
	"testing"
)

func TestCheckStudentBanStatus(t *testing.T) {

	tests := []struct {
		name    string
		email string
		want    bool
		wantErr bool
	}{
	{
			name:    "Not Banned",
			email:  "asd@gmail.com",
			want:    false,
			wantErr: false,
		},{
			name:    "Banned",
			email:  "banned@gmail.com",
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckStudentBanStatus(dbConnection,tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckStudentBanStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckStudentBanStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
