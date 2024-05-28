package utils

import (
	"reflect"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

func generateTestEmployeeJWT(employeeID string) string {
	claims := &EmployeeClaims{
		EmployeeID: employeeID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(employeeJWTKey)
	return tokenString
}

func generateTestStudentJWT(studentID string) string {
	claims := &StudentClaims{
		StudentID: studentID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(studentJWTKey)
	return tokenString
}

func generateExpiredEmployeeJWT(employeeID string) string {
	claims := &EmployeeClaims{
		EmployeeID: employeeID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(employeeJWTKey)
	return tokenString
}

func generateExpiredStudentJWT(studentID string) string {
	claims := &StudentClaims{
		StudentID: studentID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(studentJWTKey)
	return tokenString
}

func TestGenerateEmployeeJWT(t *testing.T) {
	tests := []struct {
		name      string
		employeeID string
		wantErr   bool
	}{
		{
			name:      "Valid Employee ID",
			employeeID: "employee123",
			wantErr:   false,
		},
		{
			name:      "Empty Employee ID",
			employeeID: "",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateEmployeeJWT(tt.employeeID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateEmployeeJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			claims := &EmployeeClaims{}
			_, err = jwt.ParseWithClaims(got, claims, func(token *jwt.Token) (interface{}, error) {
				return employeeJWTKey, nil
			})
			if err != nil {
				t.Errorf("Failed to parse token: %v", err)
				return
			}
			if claims.EmployeeID != tt.employeeID {
				t.Errorf("GenerateEmployeeJWT() EmployeeID = %v, want %v", claims.EmployeeID, tt.employeeID)
			}
		})
	}
}

func TestGenerateStudentJWT(t *testing.T) {
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
			got, err := GenerateStudentJWT(tt.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateStudentJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			claims := &StudentClaims{}
			_, err = jwt.ParseWithClaims(got, claims, func(token *jwt.Token) (interface{}, error) {
				return studentJWTKey, nil
			})
			if err != nil {
				t.Errorf("Failed to parse token: %v", err)
				return
			}
			if claims.StudentID != tt.studentID {
				t.Errorf("GenerateStudentJWT() StudentID = %v, want %v", claims.StudentID, tt.studentID)
			}
		})
	}
}

func TestValidateEmployeeJWT(t *testing.T) {
	tests := []struct {
		name        string
		tokenString string
		want        *EmployeeClaims
		wantErr     bool
	}{
		{
			name:        "Valid Token",
			tokenString: generateTestEmployeeJWT("employee123"),
			want: &EmployeeClaims{
				EmployeeID: "employee123",
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
				},
			},
			wantErr: false,
		},
		{
			name:        "Invalid Token",
			tokenString: "invalidTokenString",
			want:        nil,
			wantErr:     true,
		},
		{
			name:        "Expired Token",
			tokenString: generateExpiredEmployeeJWT("employee123"),
			want:        nil,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateEmployeeJWT(tt.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEmployeeJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateEmployeeJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateStudentJWT(t *testing.T) {
	tests := []struct {
		name        string
		tokenString string
		want        *StudentClaims
		wantErr     bool
	}{
		{
			name:        "Valid Token",
			tokenString: generateTestStudentJWT("student123"),
			want: &StudentClaims{
				StudentID: "student123",
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
				},
			},
			wantErr: false,
		},
		{
			name:        "Invalid Token",
			tokenString: "invalidTokenString",
			want:        nil,
			wantErr:     true,
		},
		{
			name:        "Expired Token",
			tokenString: generateExpiredStudentJWT("student123"),
			want:        nil,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateStudentJWT(tt.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateStudentJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateStudentJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}


