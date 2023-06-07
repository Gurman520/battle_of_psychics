package server

import (
	view "battle_of_psychics/backend/ViewStruct"
	"battle_of_psychics/backend/battle"
	"battle_of_psychics/backend/convertor"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

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

func (s *Server) Rank(session *sessions.Session, number int) (view.ViewDataRankPage, error) {
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

	history := convertor.ConvertRank(*s.Battles[sessionID])

	reliability := convertor.ConvertReliability(*s.Battles[sessionID])

	data := view.ViewDataRankPage{
		Title:       "Результаты битвы",
		History:     history,
		Reliability: reliability,
	}
	return data, nil
}
