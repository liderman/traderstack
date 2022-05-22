package domain

import "github.com/shopspring/decimal"

func HistoricCandleToSliceClose(candles []*HistoricCandle) []decimal.Decimal {
	var ret []decimal.Decimal
	for _, v := range candles {
		ret = append(ret, v.Close)
	}
	return ret
}
