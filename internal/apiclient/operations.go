package apiclient

import (
	"context"
	"fmt"
	"github.com/liderman/traderstack/internal/domain"
	"github.com/liderman/traderstack/internal/grpc/gen/tinkoff/investapi"
	"github.com/liderman/traderstack/internal/grpc/mapper"
)

func (a *RealClient) GetPortfolio(accountId string) ([]*domain.PortfolioPosition, error) {
	resp, err := a.cOperations.GetPortfolio(context.Background(), &investapi.PortfolioRequest{
		AccountId: accountId,
	})
	if err != nil {
		return nil, fmt.Errorf("error api GetPortfolio: %w", err)
	}

	return mapper.MapFromPortfolioPositions(resp.Positions), nil
}
