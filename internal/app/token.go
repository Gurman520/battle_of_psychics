package app

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := []byte("my_secret_key")

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	fmt.Println(tokenString)
	return tokenString, nil
}
