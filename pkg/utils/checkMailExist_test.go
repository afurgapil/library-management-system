package utils

import (
	"testing"
)

func TestCheckMailExist(t *testing.T) {
	tests := []struct {
		name    string
		mail    string
		want    bool
		wantErr bool
	}{
		{
			name:    "Mail Exist",
			mail:  "asd@gmail.com",
			want:    true,
			wantErr: false,
		},{
			name:    "Mail Not Exist",
			mail:  "a@gmail.com",
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckMailExist(dbConnection, tt.mail)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckMailExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckMailExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
