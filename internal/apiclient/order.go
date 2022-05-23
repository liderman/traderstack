package apiclient

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/liderman/traderstack/internal/domain"
	"github.com/liderman/traderstack/internal/grpc/gen/tinkoff/investapi"
	"github.com/liderman/traderstack/internal/grpc/mapper"
	"github.com/shopspring/decimal"
)

func (a *RealClient) PostOrder(figi string, lots int64, price decimal.Decimal, direction domain.OrderDirection, accountId string, orderType domain.OrderType) (*domain.PostOrderResponse, error) {
	resp, err := a.cOrder.PostOrder(context.Background(), &investapi.PostOrderRequest{
		Figi:      figi,
		Quantity:  lots,
		Price:     mapper.MapToQuotation(price),
		Direction: mapper.MapToOrderDirection(direction),
		AccountId: accountId,
		OrderType: mapper.MapToOrderType(orderType),
		OrderId:   uuid.New().String(),
	})
	if err != nil {
		return nil, fmt.Errorf("error api PostOrder: %w", err)
	}

	return mapper.MapFromPostOrderResponse(resp), nil
}

func (a *RealClient) GetOrders(accountId string) ([]*domain.OrderState, error) {
	resp, err := a.cOrder.GetOrders(context.Background(), &investapi.GetOrdersRequest{
		AccountId: accountId,
	})
	if err != nil {
		return nil, fmt.Errorf("error api GetOrders: %w", err)
	}

	return mapper.MapFromOrderStates(resp.Orders), nil
}
