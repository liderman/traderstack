package domain

import (
	"github.com/shopspring/decimal"
	"time"
)

type OrderDirection int32

const (
	OrderDirectionUnspecified OrderDirection = 0 // Значение не указано
	OrderDirectionBuy         OrderDirection = 1 // Покупка
	OrderDirectionSell        OrderDirection = 2 // Продажа
)

type OrderType int32

const (
	OrderTypeUnspecified OrderType = 0 // Значение не указано
	OrderTypeLimit       OrderType = 1 // Лимитная
	OrderTypeMarket      OrderType = 2 // Рыночная
)

// OrderExecutionReportStatus Текущий статус заявки (поручения)
type OrderExecutionReportStatus int32

const (
	OrderExecutionReportStatusUnspecified   OrderExecutionReportStatus = 0
	OrderExecutionReportStatusFill          OrderExecutionReportStatus = 1 // Исполнена
	OrderExecutionReportStatusRejected      OrderExecutionReportStatus = 2 // Отклонена
	OrderExecutionReportStatusCancelled     OrderExecutionReportStatus = 3 // Отменена пользователем
	OrderExecutionReportStatusNew           OrderExecutionReportStatus = 4 // Новая
	OrderExecutionReportStatusPartiallyfill OrderExecutionReportStatus = 5 // Частично исполнена
)

func (o OrderExecutionReportStatus) InProcess() bool {
	return o == OrderExecutionReportStatusNew || o == OrderExecutionReportStatusPartiallyfill
}

type MoneyValue struct {
	Currency string
	Value    decimal.Decimal
}

type PostOrderResponse struct {
	OrderId               string
	ExecutionReportStatus OrderExecutionReportStatus
	LotsRequested         int64
	LotsExecuted          int64
	InitialOrderPrice     *MoneyValue
	ExecutedOrderPrice    *MoneyValue
	TotalOrderAmount      *MoneyValue
	InitialCommission     *MoneyValue
	ExecutedCommission    *MoneyValue
	AciValue              *MoneyValue
	Figi                  string
	Direction             OrderDirection
	InitialSecurityPrice  *MoneyValue
	OrderType             OrderType
	Message               string
	InitialOrderPricePt   decimal.Decimal
}

// OrderState Информация о торговом поручении.
type OrderState struct {
	OrderId               string
	ExecutionReportStatus OrderExecutionReportStatus // Текущий статус заявки.
	LotsRequested         int64
	LotsExecuted          int64
	InitialOrderPrice     *MoneyValue // Начальная цена заявки. Произведение количества запрошенных лотов на цену.
	ExecutedOrderPrice    *MoneyValue // Исполненная цена заявки. Произведение средней цены покупки на количество лотов.
	TotalOrderAmount      *MoneyValue // Итоговая стоимость заявки, включающая все комиссии.
	AveragePositionPrice  *MoneyValue // Средняя цена позиции по сделке.
	InitialCommission     *MoneyValue // Начальная комиссия. Комиссия, рассчитанная на момент подачи заявки.
	ExecutedCommission    *MoneyValue // Фактическая комиссия по итогам исполнения заявки.
	Figi                  string
	Direction             OrderDirection
	InitialSecurityPrice  *MoneyValue   // Начальная цена за 1 инструмент. Для получения стоимости лота требуется умножить на лотность инструмента.
	Stages                []*OrderStage // Стадии выполнения заявки.
	ServiceCommission     *MoneyValue   // Сервисная комиссия.
	Currency              string        // Валюта заявки.
	OrderType             OrderType     // Тип заявки.
	OrderDate             time.Time     // Дата и время выставления заявки в часовом поясе UTC.
}

// OrderStage Сделки в рамках торгового поручения.
type OrderStage struct {
	Price    *MoneyValue // Цена за 1 инструмент. Для получения стоимости лота требуется умножить на лотность инструмента..
	Quantity int64       // Количество лотов.
	TradeId  string      // Идентификатор торговой операции.
}
