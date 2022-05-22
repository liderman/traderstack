package stackfuncs

import (
	"fmt"
	"github.com/liderman/traderstack/internal/algofunc"
	"github.com/liderman/traderstack/internal/datamanager"
	"github.com/liderman/traderstack/internal/domain"
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type Rsi struct {
	marketData *datamanager.MarketData
}

func NewRsi(marketData *datamanager.MarketData) *Rsi {
	return &Rsi{
		marketData: marketData,
	}
}

func (r *Rsi) Name() string {
	return "RSI"
}

func (r *Rsi) BaseType() string {
	return "decimal"
}

func (r *Rsi) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	figi, err := options.GetString("figi")
	if err != nil {
		return nil, err
	}

	period, err := options.GetInteger("period")
	if err != nil {
		return nil, err
	}

	data, err := r.marketData.GetLastCandles(figi, domain.CandleInterval1Day, int(period), now)
	if err != nil {
		return nil, fmt.Errorf("can't load last candles by figi %s and period %d: %w", figi, period, err)
	}

	rsi := algofunc.CalcRsi(domain.HistoricCandleToSliceClose(data))
	return &rsi, nil
}

func (r *Rsi) Arguments() []*engine.Argument {
	return []*engine.Argument{
		{
			Id:           "figi",
			Name:         "FIGI",
			Desc:         "Идентификатор инструмента",
			BaseType:     "string",
			ExtendedType: "figi-select",
			Required:     true,
			Value:        "",
		},
		{
			Id:           "period",
			Name:         "Период",
			Desc:         "Рекомендуется 14 дней",
			BaseType:     "integer",
			ExtendedType: "",
			Required:     true,
			Value:        int64(14),
		},
	}
}
