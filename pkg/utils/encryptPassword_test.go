package utils

import "testing"

func TestEncryptPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Encrypt password",
			args:    args{password: "asd"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncryptPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if got == "" {
				t.Errorf("EncryptPassword() returned an empty string")
				return
			}
			if got == tt.args.password {
				t.Errorf("EncryptPassword() = %v, want not equal to input password %v", got, tt.args.password)
			}
		})
	}
}
