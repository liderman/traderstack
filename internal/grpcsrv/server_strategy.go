package grpcsrv

import (
	"context"
	"github.com/liderman/traderstack/internal/engine"
	strategyv1 "github.com/liderman/traderstack/internal/grpcsrv/gen/go/liderman/traderstack/strategy/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StrategyServer struct {
	se          *engine.StrategyEngine
	mapProtobuf StrategyToProtobufMapper
	mapDomain   StrategyToDomainMapper
}

func NewStrategyServer(se *engine.StrategyEngine) *StrategyServer {
	return &StrategyServer{
		se: se,
	}
}

func (s *StrategyServer) Create(ctx context.Context, in *strategyv1.CreateRequest) (*strategyv1.CreateResponse, error) {
	strategy := s.se.Create()

	return &strategyv1.CreateResponse{
		Strategy: s.mapProtobuf.MapStrategy(strategy),
	}, nil
}

func (s *StrategyServer) Get(ctx context.Context, in *strategyv1.GetRequest) (*strategyv1.GetResponse, error) {
	strategy := s.se.Get(in.Id)

	return &strategyv1.GetResponse{
		Strategy: s.mapProtobuf.MapStrategy(strategy),
	}, nil
}

func (s *StrategyServer) GetAll(ctx context.Context, in *strategyv1.GetAllRequest) (*strategyv1.GetAllResponse, error) {
	strategies := s.se.GetAll()

	return &strategyv1.GetAllResponse{
		Strategies: s.mapProtobuf.MapStrategies(strategies),
	}, nil
}

func (s *StrategyServer) Update(ctx context.Context, in *strategyv1.UpdateRequest) (*strategyv1.UpdateResponse, error) {
	err := s.se.Update(
		s.mapDomain.MapStrategy(in.Strategy),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed update strategy: %s", err)
	}

	return &strategyv1.UpdateResponse{}, nil
}

func (s *StrategyServer) Delete(ctx context.Context, in *strategyv1.DeleteRequest) (*strategyv1.DeleteResponse, error) {
	s.se.Delete(in.Id)

	return &strategyv1.DeleteResponse{}, nil
}

func (s *StrategyServer) GetLogs(ctx context.Context, in *strategyv1.GetLogsRequest) (*strategyv1.GetLogsResponse, error) {
	logs := s.se.GetLogs(in.Id)

	return &strategyv1.GetLogsResponse{
		Logs: s.mapProtobuf.MapLogs(logs),
	}, nil
}
