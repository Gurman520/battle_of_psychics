package backend

import (
	bat "battle_of_psychics/backend/battle"

	"github.com/gorilla/sessions"
)

type Server struct {
	Battle   map[*sessions.Session]*bat.Battle
	Sessions *sessions.CookieStore
}

func CreateNewServer() *Server {
	key := []byte("super-secret-key")
	return &Server{
		Battle:   make(map[*sessions.Session]*bat.Battle),
		Sessions: sessions.NewCookieStore(key),
	}
}
