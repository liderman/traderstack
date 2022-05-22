package datamanager

import (
	"errors"
	"fmt"
	"github.com/liderman/traderstack/internal/apiclient"
	"github.com/liderman/traderstack/internal/domain"
	"github.com/patrickmn/go-cache"
	"sort"
	"time"
)

type MarketData struct {
	client apiclient.ApiClient
	cache  *cache.Cache
}

func NewMarketData(client apiclient.ApiClient, cache *cache.Cache) *MarketData {
	return &MarketData{
		client: client,
		cache:  cache,
	}
}

func (m *MarketData) GetLastPrices() ([]*domain.LastPrice, error) {
	if v, ok := m.cache.Get("lastPrices"); ok {
		return v.([]*domain.LastPrice), nil
	}

	resp, err := m.client.GetLastPrices()
	if err != nil {
		return nil, fmt.Errorf("error api GetLastPrices: %w", err)
	}

	m.cache.Set("lastPrices", resp, time.Second*2)

	return resp, nil
}

func (m *MarketData) GetLastPrice(figi string) (*domain.LastPrice, error) {
	prices, err := m.GetLastPrices()
	if err != nil {
		return nil, err
	}

	for _, v := range prices {
		if v.Figi == figi {
			return v, nil
		}
	}

	return nil, errors.New("price not found")
}

func (m *MarketData) GetCandles(figi string, from time.Time, to time.Time, interval domain.CandleInterval) ([]*domain.HistoricCandle, error) {
	cKey, cTime := m.getCandlesCacheParams(figi, from, to, interval)
	if v, ok := m.cache.Get(cKey); ok {
		return v.([]*domain.HistoricCandle), nil
	}

	resp, err := m.client.GetCandles(figi, from, to, interval)
	if err != nil {
		return nil, fmt.Errorf("error api GetCandles: %w", err)
	}

	sort.Slice(resp, func(i, j int) bool {
		return resp[i].Time.After(resp[j].Time)
	})

	m.cache.Set(cKey, resp, cTime)

	return resp, nil
}

func (m *MarketData) GetLastCandles(figi string, interval domain.CandleInterval, period int, now time.Time) ([]*domain.HistoricCandle, error) {
	limit := apiclient.LimitGetHistoricCandlesDuration(interval)
	to := now
	from := to.Add(-limit)

	var candles []*domain.HistoricCandle
	var data []*domain.HistoricCandle
	var err error
	for len(candles) < period {
		data, err = m.GetCandles(figi, from, to, interval)
		if err != nil {
			break
		}

		candles = append(candles, data...)
		to = to.Add(-limit)
		from = from.Add(-limit)
	}

	if err != nil {
		return nil, err
	}

	return candles[:period], nil
}

func (m *MarketData) getCandlesCacheParams(figi string, from time.Time, to time.Time, interval domain.CandleInterval) (string, time.Duration) {
	truncDuration := time.Second
	switch interval {
	case domain.CandleInterval1Min:
		truncDuration = time.Second * 5
	case domain.CandleInterval5Min:
		truncDuration = time.Second * 15
	case domain.CandleInterval15Min:
		truncDuration = time.Minute
	case domain.CandleInterval1Hour:
		truncDuration = time.Minute * 5
	case domain.CandleInterval1Day:
		truncDuration = time.Minute * 15
	}

	key := fmt.Sprintf(
		"%d.%s.%s.%s",
		interval,
		figi,
		from.Truncate(truncDuration),
		to.Truncate(truncDuration),
	)
	return key, truncDuration
}
