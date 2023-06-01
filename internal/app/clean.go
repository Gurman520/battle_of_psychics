package app

import (
	"battle_of_psychics/internal/def"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Удаляет не валидные сессии, чтобы не занимали место
func (svc *BattleServerSvc) CleanMap() error {
	go svc.clean()
	return nil
}

func (svc *BattleServerSvc) clean() {
	time.Sleep(time.Hour)
	svc.l.Info("Start goroutine clean not valid session")
	tokenDelete := []string{}
	for {
		for token, _ := range svc.sessions {
			if !valid(token) {
				tokenDelete = append(tokenDelete, token)
			}
		}
		for _, token := range tokenDelete {
			delete(svc.sessions, token)
		}
		svc.l.Info("Clean not valid session is = ", len(tokenDelete))
		time.Sleep(time.Hour)
	}
}

func valid(stringToken string) bool {
	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(stringToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(def.SecretKey), nil
	})

	valid := token.Claims.Valid()

	if valid == nil {
		return true
	}
	return false
}
