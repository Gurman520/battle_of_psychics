package app

import (
	"battle_of_psychics/internal/def"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := []byte(def.SecretKey)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	fmt.Println(tokenString)
	return tokenString, nil
}
