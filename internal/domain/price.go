package domain

import (
	"github.com/shopspring/decimal"
	"time"
)

// LastPrice Информация о цене.
type LastPrice struct {
	Figi  string          // Идентификатор инструмента.
	Price decimal.Decimal //Последняя цена за 1 инструмент. Для получения стоимости лота требуется умножить на лотность инструмента.
	Time  time.Time       // Время получения последней цены в часовом поясе UTC по времени биржи.
}
