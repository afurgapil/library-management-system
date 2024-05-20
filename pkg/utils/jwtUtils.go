package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var employeeJWTKey = []byte("your_secret_key")
var studentJWTKey = []byte("your_secret_key")

type EmployeeClaims struct {
	EmployeeID string `json:"employee_id"`
	jwt.StandardClaims
}

type StudentClaims struct {
	StudentID string `json:"student_id"`
	jwt.StandardClaims
}

func GenerateEmployeeJWT(employeeID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &EmployeeClaims{
		EmployeeID: employeeID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(employeeJWTKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateStudentJWT(studentID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &StudentClaims{
		StudentID: studentID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(studentJWTKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateEmployeeJWT(tokenString string) (*EmployeeClaims, error) {
	claims := &EmployeeClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return employeeJWTKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

func ValidateStudentJWT(tokenString string) (*StudentClaims, error) {
	claims := &StudentClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return studentJWTKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
