package utils

import (
	"errors"
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

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// Check if the signing method is correct

		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		// If method we used to sign the token is different from which this token has return error
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		// Return byte stringed secretKey
		return []byte(secretKey), nil
	})

	if err != nil {
		return errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return errors.New("invalid token")
	}

	// claims, ok := parsedToken.Claims.(jwt.MapClaims)

	// if !ok {
	// 	return errors.New("Invalid token claims")
	// }

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)

	return nil
}