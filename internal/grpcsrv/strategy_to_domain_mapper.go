package grpcsrv

import (
	"github.com/liderman/traderstack/internal/domain"
	strategyv1 "github.com/liderman/traderstack/internal/grpcsrv/gen/go/liderman/traderstack/strategy/v1"
)

type StrategyToDomainMapper struct {
}

func (t *StrategyToDomainMapper) MapStrategy(in *strategyv1.Strategy) *domain.Strategy {
	if in == nil {
		return nil
	}
	return &domain.Strategy{
		Id:          in.Id,
		StackId:     in.StackId,
		AccountId:   in.AccountId,
		RunInterval: in.RunIntervalDuration.AsDuration(),
		Enabled:     in.Enabled,
	}
}
