package utils

import (
	"reflect"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

func generateTestResetToken(studentID string) string {
	claims := &ResetPasswordClaims{
		StudentID: studentID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(resetPasswordJWTKey)
	return tokenString
}

func generateExpiredResetToken(studentID string) string {
	claims := &ResetPasswordClaims{
		StudentID: studentID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-1 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(resetPasswordJWTKey)
	return tokenString
}


func TestGeneratePasswordResetToken(t *testing.T) {
	tests := []struct {
		name      string
		studentID string
		wantErr   bool
	}{
		{
			name:      "Valid Student ID",
			studentID: "student123",
			wantErr:   false,
		},
		{
			name:      "Empty Student ID",
			studentID: "",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GeneratePasswordResetToken(tt.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratePasswordResetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			claims := &ResetPasswordClaims{}
			_, err = jwt.ParseWithClaims(got, claims, func(token *jwt.Token) (interface{}, error) {
				return resetPasswordJWTKey, nil
			})
			if err != nil {
				t.Errorf("Failed to parse token: %v", err)
				return
			}
			if claims.StudentID != tt.studentID {
				t.Errorf("GeneratePasswordResetToken() StudentID = %v, want %v", claims.StudentID, tt.studentID)
			}
		})
	}
}

func TestValidatePasswordResetToken(t *testing.T) {
	tests := []struct {
		name        string
		tokenString string
		want        string
		wantErr     bool
	}{
		{
			name:        "Valid Token",
			tokenString: generateTestResetToken("student123"),
			want:        "student123",
			wantErr:     false,
		},
		{
			name:        "Invalid Token",
			tokenString: "invalidTokenString",
			want:        "",
			wantErr:     true,
		},
		{
			name:        "Expired Token",
			tokenString: generateExpiredResetToken("student123"),
			want:        "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidatePasswordResetToken(tt.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePasswordResetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidatePasswordResetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

//TODO check
// func TestSendPasswordResetEmail(t *testing.T) {
// 	err := godotenv.Load("../../.env.test")
// 	if err != nil {
// 		t.Fatalf("Failed to load environment variables: %v", err)
// 	}

// 	tests := []struct {
// 		name    string
// 		email   string
// 		token   string
// 		wantErr bool
// 	}{
// 		{
// 			name:    "Valid Email and Token",
// 			email:   "test@example.com",
// 			token:   generateTestResetToken("student123"),
// 			wantErr: false,
// 		},
// 		{
// 			name:    "Invalid Email",
// 			email:   "invalid",
// 			token:   generateTestResetToken("student123"),
// 			wantErr: true,
// 		},
// 		{
// 			name:    "Empty Email",
// 			email:   "",
// 			token:   generateTestResetToken("student123"),
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := SendPasswordResetEmail(tt.email, tt.token)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("SendPasswordResetEmail() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

