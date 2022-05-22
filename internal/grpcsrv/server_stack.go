package grpcsrv

import (
	"context"
	"github.com/liderman/traderstack/internal/engine"
	stackv1 "github.com/liderman/traderstack/internal/grpcsrv/gen/go/liderman/traderstack/stack/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StackServer struct {
	sm          *engine.StackManager
	sfr         *engine.StackFuncRepository
	ts          *engine.TestStack
	mapProtobuf ToProtobufMapper
	mapDomain   ToDomainMapper
}

func NewStackServer(sm *engine.StackManager, sfr *engine.StackFuncRepository, ts *engine.TestStack) *StackServer {
	return &StackServer{
		sm:  sm,
		sfr: sfr,
		ts:  ts,
	}
}

func (s *StackServer) Create(ctx context.Context, in *stackv1.CreateRequest) (*stackv1.CreateResponse, error) {
	stack := s.sm.Create(in.Name)

	return &stackv1.CreateResponse{
		Stack: s.mapProtobuf.MapStack(stack),
	}, nil
}

func (s *StackServer) Update(ctx context.Context, in *stackv1.UpdateRequest) (*stackv1.UpdateResponse, error) {
	stack, err := s.sm.Update(in.Id, in.Name, s.mapDomain.MapSetItems(in.Items))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Failed update stack: %s", err)
	}

	return &stackv1.UpdateResponse{
		Stack: s.mapProtobuf.MapStack(stack),
	}, nil
}

func (s *StackServer) Delete(ctx context.Context, in *stackv1.DeleteRequest) (*stackv1.DeleteResponse, error) {
	s.sm.Delete(in.Id)

	return &stackv1.DeleteResponse{}, nil
}

func (s *StackServer) Get(ctx context.Context, in *stackv1.GetRequest) (*stackv1.GetResponse, error) {
	stack := s.sm.Get(in.Id)

	return &stackv1.GetResponse{
		Stack: s.mapProtobuf.MapStack(stack),
	}, nil
}

func (s *StackServer) GetAll(ctx context.Context, in *stackv1.GetAllRequest) (*stackv1.GetAllResponse, error) {
	stacks := s.sm.GetAll()

	return &stackv1.GetAllResponse{
		Stacks: s.mapProtobuf.MapStacks(stacks),
	}, nil
}

func (s *StackServer) Test(ctx context.Context, in *stackv1.TestRequest) (*stackv1.TestResponse, error) {
	resp, err := s.ts.Test(in.Id, in.Time.AsTime(), in.AccountId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Failed test stack: %s", err)
	}

	return &stackv1.TestResponse{
		Result: s.mapProtobuf.MapTestItemResults(resp),
	}, nil
}

func (s *StackServer) FuncList(ctx context.Context, in *stackv1.FuncListRequest) (*stackv1.FuncListResponse, error) {
	funcs := s.sfr.GetAllDeclaration()

	return &stackv1.FuncListResponse{
		Func: s.mapProtobuf.MapFuncs(funcs),
	}, nil
}

func (s *StackServer) FuncArgumentVarList(ctx context.Context, in *stackv1.FuncArgumentVarListRequest) (*stackv1.FuncArgumentVarListResponse, error) {
	funcList := s.sm.FuncArgumentVarList(in.StackId, in.ItemVariable, in.ArgumentId)

	return &stackv1.FuncArgumentVarListResponse{
		Variables: funcList,
	}, nil
}
