package signal

import (
	"github.com/liderman/traderstack/internal/algofunc"
	"github.com/liderman/traderstack/internal/datamanager"
	"github.com/liderman/traderstack/internal/domain"
	"github.com/liderman/traderstack/internal/engine"
	"github.com/liderman/traderstack/internal/engine/baseoption"
	"github.com/liderman/traderstack/internal/engine/option"
	"github.com/shopspring/decimal"
)

// Индекс относительной силы (RSI от англ. relative strength index
// RSI = 100 — 100/ (1+RS)
type Rsi struct {
	marketData        *datamanager.MarketData
	figi              string
	interval          domain.CandleInterval
	period            int64
	upperRsiThreshold int64
	lowerRsiThreshold int64
}

func (r *Rsi) Init(options *engine.Options) error {
	var err error
	r.figi, err = options.GetString("figi")
	if err != nil {
		return err
	}
	interval, err := options.GetInt64("interval")
	if err != nil {
		return err
	}
	r.interval = domain.CandleInterval(interval)
	r.period, err = options.GetInt64("period")
	if err != nil {
		return err
	}
	r.upperRsiThreshold, err = options.GetInt64("upperRsiThreshold")
	if err != nil {
		return err
	}
	r.lowerRsiThreshold, err = options.GetInt64("lowerRsiThreshold")
	if err != nil {
		return err
	}
	return nil
}

func (r *Rsi) Calc() *engine.Signal {
	data, err := r.marketData.GetLastCandles(r.figi, r.interval, int(r.period))
	if err != nil {
		// TODO: Добавить логирование ошибки
		return nil
	}
	rsi := algofunc.CalcRsi(domain.HistoricCandleToSliceClose(data))

	upperRsiThreshold := decimal.NewFromInt(r.upperRsiThreshold)
	lowerRsiThreshold := decimal.NewFromInt(r.lowerRsiThreshold)
	if rsi.GreaterThan(upperRsiThreshold) {
		return &engine.Signal{
			Action:  "sell",
			IsShort: true,
		}
	}

	if rsi.LessThan(lowerRsiThreshold) {
		return &engine.Signal{
			Action:  "buy",
			IsShort: true,
		}
	}

	return nil
}

func (r *Rsi) GetOptions() []engine.Option {
	var options []engine.Option
	figi := option.NewFigiByTicker("figi")
	figi.Required = true
	options = append(options, figi)

	interval := option.NewCandleInterval("interval")
	interval.Default = int64(domain.CandleInterval1Day)
	interval.Desc = "Рекомендуется, 1 день"
	options = append(options, interval)

	period := baseoption.NewInteger("period")
	period.Name = "Период"
	period.Desc = "Количество свечей для анализа, реком 14"
	minPeriod := int64(2)
	period.Min = &minPeriod
	period.Default = 14
	options = append(options, period)

	upperRsiThreshold := baseoption.NewInteger("upperRsiThreshold")
	upperRsiThreshold.Name = "Верхнее значение RSI"
	upperRsiThreshold.Desc = "Значение больше даст сигнал на покупку в short"
	upperRsiThresholdMin := int64(1)
	upperRsiThreshold.Min = &upperRsiThresholdMin
	upperRsiThresholdMax := int64(100)
	upperRsiThreshold.Max = &upperRsiThresholdMax
	upperRsiThreshold.Default = 70
	options = append(options, upperRsiThreshold)

	lowerRsiThreshold := baseoption.NewInteger("lowerRsiThreshold")
	lowerRsiThreshold.Name = "Период"
	lowerRsiThreshold.Desc = "Значение ниже даст сигнал на покупку в long"
	lowerRsiThresholdMin := int64(1)
	lowerRsiThreshold.Min = &lowerRsiThresholdMin
	lowerRsiThresholdMax := int64(100)
	upperRsiThreshold.Max = &lowerRsiThresholdMax
	lowerRsiThreshold.Default = 30
	options = append(options, lowerRsiThreshold)

	return options
}
