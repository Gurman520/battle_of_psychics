package backend

import (
	view "battle_of_psychics/backend/ViewStruct"
	"battle_of_psychics/backend/battle"
	"battle_of_psychics/backend/convertor"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func (s *Server) StartGame(w http.ResponseWriter, r *http.Request) {
	log.Println("Start Game init")
	session, _ := s.Sessions.Get(r, "cookie-name")
	session.Values["Active"] = true
	sessionID := uuid.NewString()
	session.Values["Session ID"] = sessionID
	session.Save(r, w)
	s.Battles[sessionID] = battle.NewBattle()

	data := view.ViewDataStartPage{
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
	log.Println("Hypotheses init")
	session, _ := s.Sessions.Get(r, "cookie-name")

	sessionID := session.Values["Session ID"].(string)

	if auth, ok := session.Values["Active"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Вызвать другой слой
	b := *s.Battles[sessionID]
	s.Battles[sessionID] = battle.CreateHypotheses(b)

	hypo := convertor.Convert(*s.Battles[sessionID])

	data := view.ViewDataHypothesesPage{
		Title:      "Битва экстрасенсов",
		Hypothesis: hypo,
	}
	files := []string{
		"./templates/hypotheses/hypothesesPage.tmpl",
		"./templates/base.tmpl",
		"./templates/hypotheses/hypotheses.tmpl",
		"./templates/finish.tmpl",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println("Error parser ", err)
	}
	tmpl.Execute(w, data)

}

func (s *Server) RankPsychics(w http.ResponseWriter, r *http.Request) {
	log.Println("RankPsychics init")
	session, _ := s.Sessions.Get(r, "cookie-name")

	if auth, ok := session.Values["Active"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	sessionID := session.Values["Session ID"].(string)

	r.ParseForm()
	number, _ := strconv.Atoi(r.FormValue("number"))

	// Добавить вызов следующего слоя
	s.Battles[sessionID].UserNumber = append(s.Battles[sessionID].UserNumber, number)

	b := *s.Battles[sessionID]
	s.Battles[sessionID] = battle.CalculationReliability(b)

	history := convertor.ConvertRank(*s.Battles[sessionID])

	rel := convertor.ConvertReliability(*s.Battles[sessionID])

	data := view.ViewDataRankPage{
		Title:       "Битва экстрасенсов",
		History:     history,
		Reliability: rel,
	}
	files := []string{
		"./templates/rating/ratingPage.tmpl",
		"./templates/base.tmpl",
		"./templates/rating/rating.tmpl",
		"./templates/finish.tmpl",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println("Error parser ", err)
	}
	tmpl.Execute(w, data)
}
