package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Compare the plain password with the hashed password
func CheckPassword(hashedPassword string, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
