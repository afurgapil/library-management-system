package utils

import (
	"testing"
)

func TestCheckDebit(t *testing.T) {
	tests := []struct {
		name    string
		studentID    string
		want    bool
		wantErr bool
	}{
	{
			name:    "Debit = 0",
			studentID:  "f6df1bd5-a4dc-4390-a1bc-b92d73dffec8",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Debit > 0",
			studentID:  "0d7d1132-c42c-4e64-9ff6-970310b14e28",
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckDebit(dbConnection, tt.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckDebit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckDebit() = %v, want %v", got, tt.want)
			}
		})
	}
}
