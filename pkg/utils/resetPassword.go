package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

var resetPasswordJWTKey = []byte("your_secret_key_for_reset")

type ResetPasswordClaims struct {
    StudentID string `json:"student_id"`
    jwt.StandardClaims
}

func GeneratePasswordResetToken(studentID string) (string, error) {
    expirationTime := time.Now().Add(1 * time.Hour)
    claims := &ResetPasswordClaims{
        StudentID: studentID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(resetPasswordJWTKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func ValidatePasswordResetToken(tokenString string) (string, error) {
    claims := &ResetPasswordClaims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return resetPasswordJWTKey, nil
    })

    if err != nil {
        return "", err
    }

    if !token.Valid {
        return "", jwt.ErrSignatureInvalid
    }

    return claims.StudentID, nil
}

func SendPasswordResetEmail(email, token string) error {
	 err := godotenv.Load()
    if err != nil {
        panic(err)
    }
    mail := gomail.NewMessage()
    mail.SetHeader("From", os.Getenv("MAIL_USER"))
    mail.SetHeader("To", email)
    mail.SetHeader("Subject", "Password Reset Request")
    mail.SetBody("text/html", "To reset your password, click the following link: <a href=\"https://your-domain.com/reset-password/" + token + "\">Reset Password</a>")

    dialer := gomail.NewDialer(os.Getenv("MAIL_SERVICE"), 587, os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASSWORD"))

    if err := dialer.DialAndSend(mail); err != nil {
        return err
    }

    return nil
}
