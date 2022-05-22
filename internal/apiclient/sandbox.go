package apiclient

import (
	"context"
	"fmt"
	"github.com/liderman/traderstack/internal/domain"
	"github.com/liderman/traderstack/internal/grpc/gen/tinkoff/investapi"
	"github.com/liderman/traderstack/internal/grpc/mapper"
	"github.com/shopspring/decimal"
	"time"
)

func (a *RealClient) GetSandboxAccounts() ([]*domain.Account, error) {
	resp, err := a.cSandbox.GetSandboxAccounts(context.Background(), &investapi.GetAccountsRequest{})
	if err != nil {
		return nil, fmt.Errorf("error api GetSandboxAccounts: %w", err)
	}

	return mapper.MapFromAccounts(resp.Accounts), nil
}

func (a *RealClient) OpenSandboxAccount() (string, error) {
	resp, err := a.cSandbox.OpenSandboxAccount(context.Background(), &investapi.OpenSandboxAccountRequest{})
	if err != nil {
		return "", fmt.Errorf("error api OpenSandboxAccount: %w", err)
	}

	return resp.AccountId, nil
}

func (a *RealClient) GetSandboxPositions(accountId string) (*domain.PositionsResponse, error) {
	resp, err := a.cSandbox.GetSandboxPositions(context.Background(), &investapi.PositionsRequest{
		AccountId: accountId,
	})
	if err != nil {
		return nil, fmt.Errorf("error api GetSandboxPositions: %w", err)
	}

	return mapper.MapFromPositionsResponse(resp), nil
}

func (a *RealClient) SandboxPayIn(accountId string, amount *domain.MoneyValue) error {
	_, err := a.cSandbox.SandboxPayIn(context.Background(), &investapi.SandboxPayInRequest{
		AccountId: accountId,
		Amount:    mapper.MapToMoneyValue(amount),
	})
	if err != nil {
		return fmt.Errorf("error api SandboxPayIn: %w", err)
	}

	return nil
}

func (a *RealClient) PostSandboxOrder(figi string, lots int64, price decimal.Decimal, direction domain.OrderDirection, accountId string, orderType domain.OrderType) (*domain.PostOrderResponse, error) {
	resp, err := a.cSandbox.PostSandboxOrder(context.Background(), &investapi.PostOrderRequest{
		Figi:      figi,
		Quantity:  lots,
		Price:     mapper.MapToQuotation(price),
		Direction: mapper.MapToOrderDirection(direction),
		AccountId: accountId,
		OrderType: mapper.MapToOrderType(orderType),
		OrderId:   figi + time.Now().String(),
	})
	if err != nil {
		return nil, fmt.Errorf("error api PostSandboxOrder: %w", err)
	}

	return mapper.MapFromPostOrderResponse(resp), nil
}

func (a *RealClient) GetSandboxPortfolio(accountId string) ([]*domain.PortfolioPosition, error) {
	resp, err := a.cSandbox.GetSandboxPortfolio(context.Background(), &investapi.PortfolioRequest{
		AccountId: accountId,
	})
	if err != nil {
		return nil, fmt.Errorf("error api GetSandboxPortfolio: %w", err)
	}

	return mapper.MapFromPortfolioPositions(resp.Positions), nil
}

func (a *RealClient) GetSandboxOrders(accountId string) ([]*domain.OrderState, error) {
	resp, err := a.cSandbox.GetSandboxOrders(context.Background(), &investapi.GetOrdersRequest{
		AccountId: accountId,
	})
	if err != nil {
		return nil, fmt.Errorf("error api GetSandboxOrders: %w", err)
	}

	return mapper.MapFromOrderStates(resp.Orders), nil
}
