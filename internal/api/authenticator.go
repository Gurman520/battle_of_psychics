package api

import (
	"battle_of_psychics/internal/app"
	"battle_of_psychics/openapi/models"

	"github.com/dgrijalva/jwt-go"
)

func ValidateHeader(stringToken string) (*models.Principal, error) {
	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(stringToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})

	valid := token.Claims.Valid()

	if valid == nil {
		prin := models.Principal(stringToken)
		return &prin, nil
	}
	return nil, app.ErrIncorrectApiKey
}
