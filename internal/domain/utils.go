package domain

import "github.com/shopspring/decimal"

func HistoricCandleToSliceClose(candles []*HistoricCandle) []decimal.Decimal {
	ret := make([]decimal.Decimal, 0, len(candles))
	for _, v := range candles {
		ret = append(ret, v.Close)
	}
	return ret
}
