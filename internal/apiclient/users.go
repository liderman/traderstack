package apiclient

import (
	"context"
	"fmt"
	"github.com/liderman/traderstack/internal/domain"
	"github.com/liderman/traderstack/internal/grpc/gen/tinkoff/investapi"
	"github.com/liderman/traderstack/internal/grpc/mapper"
)

func (a *RealClient) GetAccounts() ([]*domain.Account, error) {
	resp, err := a.cUsers.GetAccounts(context.Background(), &investapi.GetAccountsRequest{})
	if err != nil {
		return nil, fmt.Errorf("error api GetAccounts: %w", err)
	}

	return mapper.MapFromAccounts(resp.Accounts), nil
}
