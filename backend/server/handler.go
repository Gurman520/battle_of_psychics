// This file describes all the endpoint handlers that are implemented in this project
package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// The start page. Makes a redirect to the start of the game
func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/startGame", http.StatusSeeOther)
}

// Start of the game. The handler implements the creation of the session and returns the page.
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

// Getting hypotheses. The handler calls the next layer that implements the receipt of hypotheses of each psychic, puts the information in html and returns it to the user.
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

// Getting results. The handler receives the number entered by the user, and returns all the attempts of this user, and the validity of each psychic.
func (s *Server) ResultBattleHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("RankPsychics init")

	session, _ := s.Sessions.Get(r, "cookie-name")

	r.ParseForm()
	number, _ := strconv.Atoi(r.FormValue("number"))

	data, err := s.Result(session, number)

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
