package server

import (
	bat "battle_of_psychics/backend/battle"

	"github.com/gorilla/sessions"
)

// The basic structure. Contains: map of Battles, session.CookieStore.
type Server struct {
	Battles  map[string]*bat.Battle
	Sessions *sessions.CookieStore
}

// Creating a new server
// Called once, when the application is launched
func CreateNewServer() *Server {
	key := []byte("super-secret-key")
	return &Server{
		Battles:  make(map[string]*bat.Battle),
		Sessions: sessions.NewCookieStore(key),
	}
}
