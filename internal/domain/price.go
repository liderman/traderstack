package domain

import (
	"github.com/shopspring/decimal"
	"time"
)

// LastPrice Информация о цене.
type LastPrice struct {
	// Идентификатор инструмента.
	Figi string
	//Последняя цена за 1 инструмент. Для получения стоимости лота требуется умножить на лотность инструмента.
	Price decimal.Decimal
	// Время получения последней цены в часовом поясе UTC по времени биржи.
	Time time.Time
}
