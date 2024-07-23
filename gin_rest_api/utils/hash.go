package utils

import "golang.org/x/crypto/bcrypt"


func HashPassword(password string) (string, error) {
	/* Implement your password hashing logic here */
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}