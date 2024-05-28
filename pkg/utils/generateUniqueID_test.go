package utils

import "testing"

func TestGenerateUniqueID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Generate unique ID",
			want: "f6df1bd5-a4dc-4390-a1bc-b92d73dffec8", 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateUniqueID()
			if got == "" {
				t.Errorf("GenerateUniqueID() returned an empty string")
				return
			}
			if got == tt.want {
				t.Errorf("GenerateUniqueID() = %v, want not equal to %v", got, tt.want)
			}
		})
	}
}
