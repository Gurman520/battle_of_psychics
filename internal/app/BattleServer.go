package app

import (
	"context"

	"go.uber.org/zap"
)

type BattleServerService interface {
	CreateSession(ctx context.Context) (string, error)
	GetHypothesis(ctx context.Context, token string) (Battle, error)
	GetReliability(ctx context.Context, number int, token string) (Battle, error)
}

type BattleServerSvc struct {
	l        *zap.SugaredLogger
	sessions map[string]*Session
}

func NewBattleServerService(logger *zap.SugaredLogger) BattleServerService {
	return &BattleServerSvc{
		l:        logger,
		sessions: make(map[string]*Session),
	}
}

func (svc *BattleServerSvc) CreateSession(ctx context.Context) (string, error) {
	token, err := CreateToken()
	if err != nil {
		svc.l.Error("not create jwt token")
		return "", err
	}

	session := CreateNewSession()
	svc.sessions[token] = session

	return token, nil
}

func (svc *BattleServerSvc) GetHypothesis(ctx context.Context, token string) (Battle, error) {
	if svc.sessions[token] == nil {
		return Battle{}, ErrNotFoundSession
	}

	svc.sessions[token].Battle = *CreateHypo(svc.sessions[token].Battle)

	return svc.sessions[token].Battle, nil
}

func (svc *BattleServerSvc) GetReliability(ctx context.Context, number int, token string) (Battle, error) {
	if svc.sessions[token] == nil {
		return Battle{}, ErrNotFoundSession
	}

	svc.sessions[token].Battle = *CalculationReliability(svc.sessions[token].Battle, number)

	return svc.sessions[token].Battle, nil
}
