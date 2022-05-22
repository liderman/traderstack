package grpcsrv

import (
	"context"
	"github.com/liderman/traderstack/internal/apiclient"
	"github.com/liderman/traderstack/internal/datamanager"
	"github.com/liderman/traderstack/internal/domain"
	infov1 "github.com/liderman/traderstack/internal/grpcsrv/gen/go/liderman/traderstack/info/v1"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InfoServer struct {
	instr       *datamanager.Instruments
	api         apiclient.ApiClient
	mapProtobuf InfoToProtobufMapper
}

func NewInfoServer(instr *datamanager.Instruments, api apiclient.ApiClient) *InfoServer {
	return &InfoServer{
		instr: instr,
		api:   api,
	}
}

func (s *InfoServer) SearchInstrument(ctx context.Context, in *infov1.SearchInstrumentRequest) (*infov1.SearchInstrumentResponse, error) {
	shares, err := s.instr.SearchShares(in.Ticker, 10)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed get shares: %s", err)
	}

	return &infov1.SearchInstrumentResponse{
		Instruments: s.mapProtobuf.MapInstruments(shares),
	}, nil
}

func (s *InfoServer) Accounts(ctx context.Context, in *infov1.AccountsRequest) (*infov1.AccountsResponse, error) {
	accounts, err := s.api.GetAccounts()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed get accounts: %s", err)
	}

	return &infov1.AccountsResponse{
		Accounts: s.mapProtobuf.MapAccounts(accounts),
	}, nil
}

func (s *InfoServer) SandboxAccounts(ctx context.Context, in *infov1.SandboxAccountsRequest) (*infov1.SandboxAccountsResponse, error) {
	accounts, err := s.api.GetSandboxAccounts()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed get sandbox accounts: %s", err)
	}

	return &infov1.SandboxAccountsResponse{
		Accounts: s.mapProtobuf.MapAccounts(accounts),
	}, nil
}

func (s *InfoServer) OpenSandboxAccount(ctx context.Context, in *infov1.OpenSandboxAccountRequest) (*infov1.OpenSandboxAccountResponse, error) {
	accountId, err := s.api.OpenSandboxAccount()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed open sandbox account: %s", err)
	}

	return &infov1.OpenSandboxAccountResponse{
		AccountId: accountId,
	}, nil
}

func (s *InfoServer) GetSandboxPositions(ctx context.Context, in *infov1.GetSandboxPositionsRequest) (*infov1.GetSandboxPositionsResponse, error) {
	positions, err := s.api.GetSandboxPositions(in.AccountId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed get sandbox positions: %s", err)
	}

	return &infov1.GetSandboxPositionsResponse{
		Money: s.mapProtobuf.MapMoneys(positions.Money),
	}, nil
}

func (s *InfoServer) SandboxPayIn(ctx context.Context, in *infov1.SandboxPayInRequest) (*infov1.SandboxPayInResponse, error) {
	val, _ := decimal.NewFromString(in.Amount)
	err := s.api.SandboxPayIn(in.AccountId, &domain.MoneyValue{
		Currency: "RUB",
		Value:    val,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed pay in sandbox: %s", err)
	}

	return &infov1.SandboxPayInResponse{}, nil
}
