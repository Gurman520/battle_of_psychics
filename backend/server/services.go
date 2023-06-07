package server

import (
	view "battle_of_psychics/backend/ViewStruct"
	"battle_of_psychics/backend/battle"
	"battle_of_psychics/backend/convertor"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

// The input function receives the session, checks the validity of the session. And if everything is correct it creates a structure that will be displayed on the start page
func (s *Server) Start(session *sessions.Session) (view.ViewDataStartPage, error) {
	session.Values["TrueGameActive"] = true
	sessionID := uuid.NewString()
	session.Values["Session ID"] = sessionID
	s.Battles[sessionID] = battle.NewBattle()

	data := view.ViewDataStartPage{
		Title: "Битва экстрасенсов",
	}
	return data, nil
}

// The input function receives the session, checks the validity of the session.
// And if everything is correct, it gets a list of hypotheses that psychics have put forward. Then creates a structure that will be displayed on the hypothesis page
func (s *Server) Hypotheses(session *sessions.Session) (view.ViewDataHypothesesPage, error) {
	if auth, ok := session.Values["TrueGameActive"].(bool); !ok || !auth {
		return view.ViewDataHypothesesPage{}, ErrForbiden
	}

	sessionID := session.Values["Session ID"].(string)

	if _, ok := s.Battles[sessionID]; !ok {
		return view.ViewDataHypothesesPage{}, ErrNotValid
	}

	b := *s.Battles[sessionID]
	s.Battles[sessionID] = battle.CreateHypotheses(b)

	hypotheses := convertor.ConvertHypotheses(*s.Battles[sessionID])

	data := view.ViewDataHypothesesPage{
		Title:      "Предположения",
		Hypothesis: hypotheses,
	}
	return data, nil
}

// The input function receives the session and the number entered by the user, checks the validity of the session.
// And if everything is correct, it calculates the reliability of each psychic. Then creates a structure that will be displayed on the results page
func (s *Server) Result(session *sessions.Session, number int) (view.ViewDataRankPage, error) {
	if auth, ok := session.Values["TrueGameActive"].(bool); !ok || !auth {
		return view.ViewDataRankPage{}, ErrForbiden
	}

	sessionID := session.Values["Session ID"].(string)

	if _, ok := s.Battles[sessionID]; !ok {
		return view.ViewDataRankPage{}, ErrNotValid
	}

	s.Battles[sessionID].UserNumber = append(s.Battles[sessionID].UserNumber, number)

	b := *s.Battles[sessionID]
	s.Battles[sessionID] = battle.CalculationReliability(b)

	history := convertor.ConvertResult(*s.Battles[sessionID])

	reliability := convertor.ConvertReliability(*s.Battles[sessionID])

	data := view.ViewDataRankPage{
		Title:       "Результаты битвы",
		History:     history,
		Reliability: reliability,
	}
	return data, nil
}
