package domain

import "github.com/shopspring/decimal"

type Share struct {
	Figi   string
	Ticker string
	// Класс-код (секция торгов).
	ClassCode string
	Isin      string
	Lot       int32
	Currency  string
	Name      string
	// Торговая площадка.
	Exchange string
	// Текущий режим торгов инструмента.
	TradingStatus SecurityTradingStatus
	// Признак внебиржевой ценной бумаги.
	OtcFlag bool
	// Признак доступности для покупки.
	BuyAvailableFlag bool
	// Признак доступности для продажи.
	SellAvailableFlag bool
	// Шаг цены.
	MinPriceIncrement decimal.Decimal
	// Признак доступности торгов через API.
	ApiTradeAvailableFlag bool
	// Уникальный идентификатор инструмента.
	Uid string
}
