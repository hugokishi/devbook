package security

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash - Function for hash user password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword - Verify the user password
func VerifyPassword(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
