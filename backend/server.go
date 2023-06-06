package backend

import (
	bat "battle_of_psychics/backend/battle"

	"github.com/gorilla/sessions"
)

type Server struct {
	Battles  map[string]*bat.Battle
	Sessions *sessions.CookieStore
}

func CreateNewServer() *Server {
	key := []byte("super-secret-key")
	return &Server{
		Battles:  make(map[string]*bat.Battle),
		Sessions: sessions.NewCookieStore(key),
	}
}
