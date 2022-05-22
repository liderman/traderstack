package apiclient

import (
	"context"
	"fmt"
	"github.com/liderman/traderstack/internal/domain"
	"github.com/liderman/traderstack/internal/grpc/gen/tinkoff/investapi"
	"github.com/liderman/traderstack/internal/grpc/mapper"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (a *RealClient) GetCandles(figi string, from time.Time, to time.Time, interval domain.CandleInterval) ([]*domain.HistoricCandle, error) {
	resp, err := a.cMarket.GetCandles(context.Background(), &investapi.GetCandlesRequest{
		Figi:     figi,
		From:     timestamppb.New(from),
		To:       timestamppb.New(to),
		Interval: mapper.MapToCandleInterval(interval),
	})
	if err != nil {
		return nil, fmt.Errorf("error api GetCandles: %w", err)
	}

	return mapper.MapFromHistoricCandles(resp.Candles), nil
}

func (a *RealClient) GetLastPrices() ([]*domain.LastPrice, error) {
	resp, err := a.cMarket.GetLastPrices(context.Background(), &investapi.GetLastPricesRequest{})
	if err != nil {
		return nil, fmt.Errorf("error api GetLastPrices: %w", err)
	}

	return mapper.MapFromLastPrices(resp.LastPrices), nil
}
