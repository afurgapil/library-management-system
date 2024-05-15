package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
    EmployeeID string `json:"employee_id"`
    jwt.StandardClaims
}

func GenerateJWT(employeeID string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        EmployeeID: employeeID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, jwt.ErrSignatureInvalid
    }

    return claims, nil
}
