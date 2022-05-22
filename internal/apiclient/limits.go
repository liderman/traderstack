package apiclient

import (
	"github.com/liderman/traderstack/internal/domain"
	"time"
)

func LimitGetHistoricCandlesDuration(interval domain.CandleInterval) time.Duration {
	switch interval {
	case domain.CandleInterval1Min:
		return time.Hour * 24
	case domain.CandleInterval5Min:
		return time.Hour * 24
	case domain.CandleInterval15Min:
		return time.Hour * 24
	case domain.CandleInterval1Hour:
		return time.Hour * 24 * 7
	case domain.CandleInterval1Day:
		return time.Hour * 24 * 356
	}

	return time.Hour * 24
}
