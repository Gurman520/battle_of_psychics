package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/startGame", http.StatusSeeOther)
}

func (s *Server) StartHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Start Game init")
	session, _ := s.Sessions.Get(r, "cookie-name")

	data, err := s.Start(session)

	switch err {
	case nil:
		session.Save(r, w)
		tmpl, err := template.ParseFiles(filesStart...)
		if err != nil {
			fmt.Println("Error parser ", err)
		}
		tmpl.Execute(w, data)
		return
	default:
		http.Error(w, "Eternal server error", http.StatusInternalServerError)
		return
	}

}

func (s *Server) HypothesesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hypotheses init")
	session, _ := s.Sessions.Get(r, "cookie-name")

	data, err := s.Hypotheses(session)

	switch err {
	case nil:
		tmpl, err := template.ParseFiles(filesHypotheses...)
		if err != nil {
			fmt.Println("Error parser ", err)
		}
		tmpl.Execute(w, data)
		return
	case ErrForbiden:
		log.Println("Hypotheses - Forbidden")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	case ErrNotValid:
		log.Println("Hypotheses ", ErrNotValid)
		http.Error(w, ErrNotValid.Error(), http.StatusBadRequest)
		return
	default:
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

}

func (s *Server) RankPsychicsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("RankPsychics init")

	session, _ := s.Sessions.Get(r, "cookie-name")

	r.ParseForm()
	number, _ := strconv.Atoi(r.FormValue("number"))

	data, err := s.Rank(session, number)

	switch err {
	case nil:
		tmpl, err := template.ParseFiles(filesRank...)
		if err != nil {
			fmt.Println("Error parser ", err)
		}
		tmpl.Execute(w, data)
		return
	case ErrForbiden:
		log.Println("RankPsychics - Forbidden")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	case ErrNotValid:
		log.Println("Rank ", ErrNotValid)
		http.Error(w, ErrNotValid.Error(), http.StatusBadRequest)
		return
	default:
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
}
