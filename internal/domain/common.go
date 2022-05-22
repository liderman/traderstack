package domain

type SecurityTradingStatus int32

const (
	// SecurityTradingStatusUnspecified Торговый статус не определён
	SecurityTradingStatusUnspecified SecurityTradingStatus = 0
	// SecurityTradingStatusNotAvailableForTrading Недоступен для торгов
	SecurityTradingStatusNotAvailableForTrading SecurityTradingStatus = 1
	// SecurityTradingStatusOpeningPeriod Период открытия торгов
	SecurityTradingStatusOpeningPeriod SecurityTradingStatus = 2
	// SecurityTradingStatusClosingPeriod Период закрытия торгов
	SecurityTradingStatusClosingPeriod SecurityTradingStatus = 3
	// SecurityTradingStatusBreakInTrading Перерыв в торговле
	SecurityTradingStatusBreakInTrading SecurityTradingStatus = 4
	// SecurityTradingStatusNormalTrading Нормальная торговля
	SecurityTradingStatusNormalTrading SecurityTradingStatus = 5
	// SecurityTradingStatusClosingAuction Аукцион закрытия
	SecurityTradingStatusClosingAuction SecurityTradingStatus = 6
	// SecurityTradingStatusDarkPoolAuction Аукцион крупных пакетов
	SecurityTradingStatusDarkPoolAuction SecurityTradingStatus = 7
	// SecurityTradingStatusDiscreteAuction Дискретный аукцион
	SecurityTradingStatusDiscreteAuction SecurityTradingStatus = 8
	// SecurityTradingStatusOpeningAuctionPeriod Аукцион открытия
	SecurityTradingStatusOpeningAuctionPeriod SecurityTradingStatus = 9
	// SecurityTradingStatusTradingAtClosingAuctionPrice Период торгов по цене аукциона закрытия
	SecurityTradingStatusTradingAtClosingAuctionPrice SecurityTradingStatus = 10
	// SecurityTradingStatusSessionAssigned Сессия назначена
	SecurityTradingStatusSessionAssigned SecurityTradingStatus = 11
	// SecurityTradingStatusSessionClose Сессия закрыта
	SecurityTradingStatusSessionClose SecurityTradingStatus = 12
	// SecurityTradingStatusSessionOpen Сессия открыта
	SecurityTradingStatusSessionOpen SecurityTradingStatus = 13
	// SecurityTradingStatusDealerNormalTrading Доступна торговля в режиме внутренней ликвидности брокера
	SecurityTradingStatusDealerNormalTrading SecurityTradingStatus = 14
	// SecurityTradingStatusDealerBreakInTrading Перерыв торговли в режиме внутренней ликвидности брокера
	SecurityTradingStatusDealerBreakInTrading SecurityTradingStatus = 15
	// SecurityTradingStatusDealerNotAvailableForTrading Недоступна торговля в режиме внутренней ликвидности брокера
	SecurityTradingStatusDealerNotAvailableForTrading SecurityTradingStatus = 16
)

// AllowLimitOrder Возможность выставлять лимитные заявки
func (s SecurityTradingStatus) AllowLimitOrder() bool {
	switch s {
	case SecurityTradingStatusOpeningPeriod, SecurityTradingStatusClosingPeriod, SecurityTradingStatusNormalTrading,
		SecurityTradingStatusClosingAuction, SecurityTradingStatusDiscreteAuction,
		SecurityTradingStatusOpeningAuctionPeriod, SecurityTradingStatusTradingAtClosingAuctionPrice,
		SecurityTradingStatusSessionAssigned, SecurityTradingStatusSessionClose, SecurityTradingStatusSessionOpen:
		return true
	}
	return false
}

// AllowMarketOrder Возможность выставлять рыночные заявки
func (s SecurityTradingStatus) AllowMarketOrder() bool {
	switch s {
	case SecurityTradingStatusNormalTrading, SecurityTradingStatusDarkPoolAuction,
		SecurityTradingStatusDealerNormalTrading:
		return true
	}
	return false
}
