package utils

import "testing"

func TestCheckMailValid(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid Mail",
			args: args{
				email: "validmail@mail.com",
			},
			want: true,
		},
		{
			name: "InvalidMail",
			args: args{
				email: "invalidmail",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckMailValid(tt.args.email); got != tt.want {
				t.Errorf("CheckMailValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
