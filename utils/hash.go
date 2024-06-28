package utils

import "golang.org/x/crypto/bcrypt"

const hashingCost = 14

// function to generate string
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hashingCost)
	return string(bytes), err
}

// function to compare/decrypt the password
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
