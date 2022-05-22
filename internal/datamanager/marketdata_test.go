package datamanager

import (
	"github.com/golang/mock/gomock"
	"github.com/liderman/traderstack/internal/apiclient"
	"github.com/liderman/traderstack/internal/apiclient/mock"
	"github.com/liderman/traderstack/internal/domain"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMarketData_GetCandles(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	client := mock.NewMockApiClient(ctl)

	figi := "test"
	to := time.Now()
	from := to.Add(-apiclient.LimitGetHistoricCandlesDuration(domain.CandleInterval1Day))
	interval := domain.CandleInterval1Day

	gomock.InOrder(
		client.EXPECT().GetCandles(figi, from, to, interval).Return([]*domain.HistoricCandle{
			{Time: time.Now()},
			{Time: time.Now().Add(-time.Hour * 24)},
			{Time: time.Now().Add(-time.Hour * 24 * 10)},
			{Time: time.Now().Add(-time.Hour * 24 * 2)},
		}, nil),
	)

	md := NewMarketData(client, cache.New(time.Minute, time.Minute))
	resp, err := md.GetCandles(figi, from, to, interval)
	assert.NoError(t, err)
	for i := 1; i < 4; i++ {
		assert.True(t, resp[i-1].Time.After(resp[i].Time))
	}
}

func TestMarketData_GetLastCandles(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	client := mock.NewMockApiClient(ctl)

	figi := "test"
	interval := domain.CandleInterval1Day

	gomock.InOrder(
		client.EXPECT().GetCandles(figi, gomock.Any(), gomock.Any(), interval).Return([]*domain.HistoricCandle{
			{Time: time.Now()},
			{Time: time.Now().Add(-time.Hour * 24)},
			{Time: time.Now().Add(-time.Hour * 24 * 3)},
			{Time: time.Now().Add(-time.Hour * 24 * 2)},
		}, nil),
		client.EXPECT().GetCandles(figi, gomock.Any(), gomock.Any(), interval).Return([]*domain.HistoricCandle{
			{Time: time.Now().Add(-time.Hour * 24 * 5)},
			{Time: time.Now().Add(-time.Hour * 24 * 4)},
			{Time: time.Now().Add(-time.Hour * 24 * 8)},
			{Time: time.Now().Add(-time.Hour * 24 * 7)},
			{Time: time.Now().Add(-time.Hour * 24 * 6)},
		}, nil),
	)

	md := NewMarketData(client, cache.New(time.Minute, time.Minute))
	resp, err := md.GetLastCandles(figi, interval, 6, time.Now())
	assert.NoError(t, err)
	assert.Len(t, resp, 6)
	for i := 1; i < 6; i++ {
		assert.True(t, resp[i-1].Time.After(resp[i].Time))
	}
}
