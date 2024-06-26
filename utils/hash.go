package utils

import "golang.org/x/crypto/bcrypt"

const hashingCost = 14

// function to generate string
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hashingCost)
	return string(bytes), err
}
