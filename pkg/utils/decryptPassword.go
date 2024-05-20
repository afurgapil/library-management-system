package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func DecryptPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}