package utils

import "golang.org/x/crypto/bcrypt"


func HashPassword(password string) (string, error) {
	/* Implement your password hashing logic here */
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hashedPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	// If err is nil, means password is valid otherwise we get false
	return err == nil
}