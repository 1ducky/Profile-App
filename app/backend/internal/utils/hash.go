package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword converts a plain text password string into a secure bcrypt hash string
func HashString(password string) (string, error) {
	// DefaultCost is currently 10, which provides an ideal speed/security balance
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash compares a raw password string against its hashed version
func CheckHashString(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
