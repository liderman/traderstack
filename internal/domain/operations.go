package domain

import (
	"github.com/shopspring/decimal"
)

// PositionsResponse Список позиций по счёту.
type PositionsResponse struct {
	// Массив валютных позиций портфеля.
	Money []*MoneyValue
	// Массив заблокированных валютных позиций портфеля.
	Blocked []*MoneyValue
	// Список ценно-бумажных позиций портфеля.
	Securities []*PositionsSecurities
	// Признак идущей в данный момент выгрузки лимитов.
	LimitsLoadingInProgress bool
}

// PositionsSecurities Баланс позиции ценной бумаги.
type PositionsSecurities struct {
	// Figi-идентификатор бумаги.
	Figi string
	// Заблокировано.
	Blocked int64
	// Текущий незаблокированный баланс.
	Balance int64
}

// PortfolioPosition Позиции портфеля.
type PortfolioPosition struct {
	Figi                     string
	InstrumentType           string
	Quantity                 decimal.Decimal // Количество инструмента в портфеле в штуках.
	AveragePositionPrice     *MoneyValue     // Средневзвешенная цена позиции. **Возможна задержка до секунды для пересчёта**.
	ExpectedYield            decimal.Decimal // Текущая рассчитанная относительная доходность позиции, в %.
	CurrentNkd               *MoneyValue     // Текущий НКД.
	AveragePositionPricePt   decimal.Decimal // Средняя цена лота в позиции в пунктах (для фьючерсов). **Возможна задержка до секунды для пересчёта**.
	CurrentPrice             *MoneyValue     // Текущая цена за 1 инструмент. Для получения стоимости лота требуется умножить на лотность инструмента.
	AveragePositionPriceFifo *MoneyValue     // Средняя цена лота в позиции по методу FIFO. **Возможна задержка до секунды для пересчёта**.
	QuantityLots             decimal.Decimal // Количество лотов в портфеле.
}
