package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns a hashed password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %s", err)
	}

	return string(hashedPassword), nil
}

// CheckPassword check if the provided password is valid
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
