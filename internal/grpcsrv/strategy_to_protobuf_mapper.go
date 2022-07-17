package grpcsrv

import (
	"github.com/liderman/traderstack/internal/domain"
	strategyv1 "github.com/liderman/traderstack/internal/grpcsrv/gen/go/liderman/traderstack/strategy/v1"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type StrategyToProtobufMapper struct {
}

func (t *StrategyToProtobufMapper) MapStrategy(in *domain.Strategy) *strategyv1.Strategy {
	if in == nil {
		return nil
	}
	return &strategyv1.Strategy{
		Id:                  in.Id,
		StackId:             in.StackId,
		AccountId:           in.AccountId,
		RunIntervalDuration: durationpb.New(in.RunInterval),
		Enabled:             in.Enabled,
	}
}

func (t *StrategyToProtobufMapper) MapLog(in *domain.StrategyLog) *strategyv1.StrategyLog {
	if in == nil {
		return nil
	}
	return &strategyv1.StrategyLog{
		LogType: string(in.LogType),
		Message: in.Message,
		Time:    timestamppb.New(in.Time),
	}
}

func (t *StrategyToProtobufMapper) MapStrategies(in []*domain.Strategy) []*strategyv1.Strategy {
	ret := make([]*strategyv1.Strategy, 0, len(in))
	for _, v := range in {
		ret = append(ret, t.MapStrategy(v))
	}

	return ret
}

func (t *StrategyToProtobufMapper) MapLogs(in []*domain.StrategyLog) []*strategyv1.StrategyLog {
	ret := make([]*strategyv1.StrategyLog, 0, len(in))
	for _, v := range in {
		ret = append(ret, t.MapLog(v))
	}

	return ret
}
