package utils

import "testing"

func TestDecryptPassword(t *testing.T) {
	type args struct {
		hashedPassword string
		password       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "True Pass",
			args: args{
				hashedPassword: "$2a$10$o7UtYevkEzR/7FsB8UXbV.58jjKA2lgC8wJFFC9qiS8dgjcVFEWQq",
				password:       "asd",
			},
			wantErr: false,
		},
		{
			name: "Wrong Pass",
			args: args{
				hashedPassword: "aasd",
				password:       "a",
			},
			wantErr: true, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DecryptPassword(tt.args.hashedPassword, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("DecryptPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
