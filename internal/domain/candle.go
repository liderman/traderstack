package domain

import (
	"github.com/shopspring/decimal"
	"time"
)

type CandleInterval int

const (
	CandleIntervalUnspecified CandleInterval = iota
	CandleInterval1Min
	CandleInterval5Min
	CandleInterval15Min
	CandleInterval1Hour
	CandleInterval1Day
)

func (c *CandleInterval) ToDuration() time.Duration {
	switch *c {
	case CandleInterval1Min:
		return time.Minute
	case CandleInterval5Min:
		return time.Minute * 5
	case CandleInterval15Min:
		return time.Minute * 15
	case CandleInterval1Hour:
		return time.Hour
	}

	return time.Nanosecond
}

type HistoricCandle struct {
	Open       decimal.Decimal
	High       decimal.Decimal
	Low        decimal.Decimal
	Close      decimal.Decimal
	Volume     int64
	Time       time.Time
	IsComplete bool
}
