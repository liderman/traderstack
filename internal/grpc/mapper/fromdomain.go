package mapper

import (
	"github.com/liderman/traderstack/internal/domain"
	"github.com/liderman/traderstack/internal/grpc/gen/tinkoff/investapi"
	"github.com/shopspring/decimal"
)

func MapToCandleInterval(in domain.CandleInterval) investapi.CandleInterval {
	switch in {
	case domain.CandleInterval1Min:
		return investapi.CandleInterval_CANDLE_INTERVAL_1_MIN
	case domain.CandleInterval5Min:
		return investapi.CandleInterval_CANDLE_INTERVAL_5_MIN
	case domain.CandleInterval15Min:
		return investapi.CandleInterval_CANDLE_INTERVAL_15_MIN
	case domain.CandleInterval1Hour:
		return investapi.CandleInterval_CANDLE_INTERVAL_HOUR
	case domain.CandleInterval1Day:
		return investapi.CandleInterval_CANDLE_INTERVAL_DAY
	}

	return investapi.CandleInterval_CANDLE_INTERVAL_UNSPECIFIED
}

func MapToQuotation(in decimal.Decimal) *investapi.Quotation {
	v := in.Sub(decimal.NewFromInt(in.IntPart())).CoefficientInt64()
	return &investapi.Quotation{
		Units: in.IntPart(),
		Nano:  int32(calcFactor(v, -9-in.Exponent())),
	}
}

func calcFactor(v int64, exp int32) int64 {
	if exp == 0 {
		return v
	}
	return calcFactor(v*10, exp+1)
}

func MapToOrderDirection(in domain.OrderDirection) investapi.OrderDirection {
	switch in {
	case domain.OrderDirectionBuy:
		return investapi.OrderDirection_ORDER_DIRECTION_BUY
	case domain.OrderDirectionSell:
		return investapi.OrderDirection_ORDER_DIRECTION_SELL
	}
	return investapi.OrderDirection_ORDER_DIRECTION_UNSPECIFIED
}

func MapToOrderType(in domain.OrderType) investapi.OrderType {
	switch in {
	case domain.OrderTypeLimit:
		return investapi.OrderType_ORDER_TYPE_LIMIT
	case domain.OrderTypeMarket:
		return investapi.OrderType_ORDER_TYPE_MARKET
	}
	return investapi.OrderType_ORDER_TYPE_UNSPECIFIED
}

func MapToMoneyValue(in *domain.MoneyValue) *investapi.MoneyValue {
	v := in.Value.Sub(decimal.NewFromInt(in.Value.IntPart())).CoefficientInt64()
	return &investapi.MoneyValue{
		Currency: in.Currency,
		Units:    in.Value.IntPart(),
		Nano:     int32(calcFactor(v, -9-in.Value.Exponent())),
	}
}
