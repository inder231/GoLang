package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "qwerty1234ytrewq"
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email": email,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})
	// To sent a string to user we using SignedString method
	// We need to provide a key to sign the token in byte slice form
	return token.SignedString([]byte(secretKey))
}