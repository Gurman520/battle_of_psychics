package app

import (
	"context"

	"go.uber.org/zap"
)

type BattleServerService interface {
	GetHypothesis(ctx context.Context) (Battle, error)
	GetReliability(ctx context.Context) (Battle, error)
}

type BattleServerSvc struct {
	l *zap.SugaredLogger
}

func NewBattleServerService(logger *zap.SugaredLogger) BattleServerService {
	return &BattleServerSvc{
		l: logger,
	}
}

func (svc *BattleServerSvc) GetHypothesis(ctx context.Context) (Battle, error) {
	return Battle{}, nil
}

func (svc *BattleServerSvc) GetReliability(ctx context.Context) (Battle, error) {
	return Battle{}, nil
}
