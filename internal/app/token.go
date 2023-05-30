package app

import (
	"battle_of_psychics/internal/def"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Creating a new token for a new session
// The token is valid for 24 hours
func CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := []byte(def.SecretKey)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
