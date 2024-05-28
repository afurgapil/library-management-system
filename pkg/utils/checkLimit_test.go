package utils

import (
	"testing"
)

func TestCheckLimit(t *testing.T) {
	
	tests := []struct {
		name    string
		studentID    string
		want    bool
		wantErr bool
	}{
		{
			name:    "Limit okay",
			studentID:  "f6df1bd5-a4dc-4390-a1bc-b92d73dffec8",
			want:    true,
			wantErr: false,
		},{
			name:    "No Limit",
			studentID:  "c96d572b-1f54-438a-8a00-6191f09645d9",
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckLimit(dbConnection, tt.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}
