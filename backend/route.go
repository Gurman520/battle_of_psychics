package backend

import (
	"battle_of_psychics/backend/battle"
	"fmt"
	"html/template"
	"net/http"
)

type ViewData struct {
	Title   string
	Message string
}

func (s *Server) StartGame(w http.ResponseWriter, r *http.Request) {
	session, _ := s.Sessions.Get(r, "cookie-name")
	session.Values["Active"] = true
	session.Save(r, w)
	s.Battle[session] = battle.NewBattle()

	data := ViewData{
		Title: "Битва экстрасенсов",
	}
	files := []string{
		"./templates/start/startPage.tmpl",
		"./templates/base.tmpl",
		"./templates/start/start.tmpl",
		"./templates/finish.tmpl",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println("Error parser ", err)
	}
	tmpl.Execute(w, data)
}

func (s *Server) Hypotheses(w http.ResponseWriter, r *http.Request) {
	session, _ := s.Sessions.Get(r, "cookie-name")

	if auth, ok := session.Values["Active"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintln(w, "The cake is a lie!")

	// Вызов метода создания гипотез.

}

func (s *Server) RankPsychics(w http.ResponseWriter, r *http.Request) {
	session, _ := s.Sessions.Get(r, "cookie-name")

	if auth, ok := session.Values["Active"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintln(w, "The cake is a lie!")

	// Вызов метода расчета достоверности экстрасенсов
}
