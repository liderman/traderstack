package mapper

import (
	"github.com/liderman/traderstack/internal/domain"
	"github.com/liderman/traderstack/internal/grpc/gen/tinkoff/investapi"
	"github.com/shopspring/decimal"
)

func MapFromHistoricCandles(in []*investapi.HistoricCandle) []*domain.HistoricCandle {
	ret := make([]*domain.HistoricCandle, 0, len(in))
	for _, v := range in {
		ret = append(ret, MapFromHistoricCandle(v))
	}
	return ret
}

func MapFromHistoricCandle(in *investapi.HistoricCandle) *domain.HistoricCandle {
	return &domain.HistoricCandle{
		Open:       MapFromQuotation(in.Open),
		High:       MapFromQuotation(in.High),
		Low:        MapFromQuotation(in.Low),
		Close:      MapFromQuotation(in.Close),
		Volume:     in.Volume,
		Time:       in.Time.AsTime(),
		IsComplete: in.IsComplete,
	}
}

func MapFromQuotation(in *investapi.Quotation) decimal.Decimal {
	if in == nil {
		return decimal.Zero
	}

	if in.Units == 0 && in.Nano == 0 {
		return decimal.Zero
	}

	return decimal.NewFromInt(in.Units).Add(
		decimal.New(int64(in.Nano), -9),
	)
}

func MapFromShares(in []*investapi.Share) []*domain.Share {
	ret := make([]*domain.Share, 0, len(in))
	for _, v := range in {
		ret = append(ret, MapFromShare(v))
	}

	return ret
}

func MapFromShare(in *investapi.Share) *domain.Share {
	return &domain.Share{
		Figi:                  in.Figi,
		Ticker:                in.Ticker,
		ClassCode:             in.ClassCode,
		Isin:                  in.Isin,
		Lot:                   in.Lot,
		Currency:              in.Currency,
		Name:                  in.Name,
		Exchange:              in.Exchange,
		TradingStatus:         MapFromSecurityTradingStatus(in.TradingStatus),
		OtcFlag:               in.OtcFlag,
		BuyAvailableFlag:      in.BuyAvailableFlag,
		SellAvailableFlag:     in.SellAvailableFlag,
		MinPriceIncrement:     MapFromQuotation(in.MinPriceIncrement),
		ApiTradeAvailableFlag: in.ApiTradeAvailableFlag,
		Uid:                   in.Uid,
	}
}

func MapFromSecurityTradingStatus(in investapi.SecurityTradingStatus) domain.SecurityTradingStatus {
	switch in {
	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_NOT_AVAILABLE_FOR_TRADING:
		return domain.SecurityTradingStatusNotAvailableForTrading
	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_OPENING_PERIOD:
		return domain.SecurityTradingStatusOpeningPeriod
	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_CLOSING_PERIOD:
		return domain.SecurityTradingStatusClosingPeriod
	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_BREAK_IN_TRADING:
		return domain.SecurityTradingStatusBreakInTrading
	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_NORMAL_TRADING:
		return domain.SecurityTradingStatusNormalTrading
	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_CLOSING_AUCTION:
		return domain.SecurityTradingStatusClosingAuction
	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_DARK_POOL_AUCTION:
		return domain.SecurityTradingStatusDarkPoolAuction
	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_DISCRETE_AUCTION:
		return domain.SecurityTradingStatusDiscreteAuction
	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_OPENING_AUCTION_PERIOD:
		return domain.SecurityTradingStatusOpeningAuctionPeriod
	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_TRADING_AT_CLOSING_AUCTION_PRICE:
		return domain.SecurityTradingStatusTradingAtClosingAuctionPrice
	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_SESSION_ASSIGNED:
		return domain.SecurityTradingStatusSessionAssigned
	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_SESSION_CLOSE:
		return domain.SecurityTradingStatusSessionClose
	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_SESSION_OPEN:
		return domain.SecurityTradingStatusSessionOpen
	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_DEALER_NORMAL_TRADING:
		return domain.SecurityTradingStatusDealerNormalTrading
	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_DEALER_BREAK_IN_TRADING:
		return domain.SecurityTradingStatusDealerBreakInTrading
	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_DEALER_NOT_AVAILABLE_FOR_TRADING:
		return domain.SecurityTradingStatusDealerNotAvailableForTrading
	}

	return domain.SecurityTradingStatusUnspecified
}

func MapFromPostOrderResponse(in *investapi.PostOrderResponse) *domain.PostOrderResponse {
	return &domain.PostOrderResponse{
		OrderId:               in.OrderId,
		ExecutionReportStatus: MapFromOrderExecutionReportStatus(in.ExecutionReportStatus),
		LotsRequested:         in.LotsRequested,
		LotsExecuted:          in.LotsExecuted,
		InitialOrderPrice:     MapFromMoneyValue(in.InitialOrderPrice),
		ExecutedOrderPrice:    MapFromMoneyValue(in.ExecutedOrderPrice),
		TotalOrderAmount:      MapFromMoneyValue(in.TotalOrderAmount),
		InitialCommission:     MapFromMoneyValue(in.InitialCommission),
		ExecutedCommission:    MapFromMoneyValue(in.ExecutedCommission),
		AciValue:              MapFromMoneyValue(in.AciValue),
		Figi:                  in.Figi,
		Direction:             MapFromOrderDirection(in.Direction),
		InitialSecurityPrice:  MapFromMoneyValue(in.InitialSecurityPrice),
		OrderType:             MapFromOrderType(in.OrderType),
		Message:               in.Message,
		InitialOrderPricePt:   MapFromQuotation(in.InitialOrderPricePt),
	}
}

func MapFromPortfolioPosition(in *investapi.PortfolioPosition) *domain.PortfolioPosition {
	return &domain.PortfolioPosition{
		Figi:                     in.Figi,
		InstrumentType:           in.InstrumentType,
		Quantity:                 MapFromQuotation(in.Quantity),
		AveragePositionPrice:     MapFromMoneyValue(in.AveragePositionPrice),
		ExpectedYield:            MapFromQuotation(in.ExpectedYield),
		CurrentNkd:               MapFromMoneyValue(in.CurrentNkd),
		AveragePositionPricePt:   MapFromQuotation(in.AveragePositionPricePt),
		CurrentPrice:             MapFromMoneyValue(in.CurrentPrice),
		AveragePositionPriceFifo: MapFromMoneyValue(in.AveragePositionPriceFifo),
		QuantityLots:             MapFromQuotation(in.QuantityLots),
	}
}

func MapFromPortfolioPositions(in []*investapi.PortfolioPosition) []*domain.PortfolioPosition {
	ret := make([]*domain.PortfolioPosition, 0, len(in))
	for _, v := range in {
		ret = append(ret, MapFromPortfolioPosition(v))
	}
	return ret
}

func MapFromOrderExecutionReportStatus(in investapi.OrderExecutionReportStatus) domain.OrderExecutionReportStatus {
	switch in {
	case investapi.OrderExecutionReportStatus_EXECUTION_REPORT_STATUS_FILL:
		return domain.OrderExecutionReportStatusFill
	case investapi.OrderExecutionReportStatus_EXECUTION_REPORT_STATUS_REJECTED:
		return domain.OrderExecutionReportStatusRejected
	case investapi.OrderExecutionReportStatus_EXECUTION_REPORT_STATUS_CANCELLED:
		return domain.OrderExecutionReportStatusCancelled
	case investapi.OrderExecutionReportStatus_EXECUTION_REPORT_STATUS_NEW:
		return domain.OrderExecutionReportStatusNew
	case investapi.OrderExecutionReportStatus_EXECUTION_REPORT_STATUS_PARTIALLYFILL:
		return domain.OrderExecutionReportStatusPartiallyfill
	}
	return domain.OrderExecutionReportStatusUnspecified
}

func MapFromOrderDirection(in investapi.OrderDirection) domain.OrderDirection {
	switch in {
	case investapi.OrderDirection_ORDER_DIRECTION_BUY:
		return domain.OrderDirectionBuy
	case investapi.OrderDirection_ORDER_DIRECTION_SELL:
		return domain.OrderDirectionSell
	}
	return domain.OrderDirectionUnspecified
}

func MapFromOrderType(in investapi.OrderType) domain.OrderType {
	switch in {
	case investapi.OrderType_ORDER_TYPE_LIMIT:
		return domain.OrderTypeLimit
	case investapi.OrderType_ORDER_TYPE_MARKET:
		return domain.OrderTypeMarket
	}
	return domain.OrderTypeUnspecified
}

func MapFromMoneyValue(in *investapi.MoneyValue) *domain.MoneyValue {
	if in == nil {
		return nil
	}

	return &domain.MoneyValue{
		Currency: in.Currency,
		Value: decimal.NewFromInt(in.Units).Add(
			decimal.New(int64(in.Nano), -9),
		),
	}
}

func MapFromMoneyValues(in []*investapi.MoneyValue) []*domain.MoneyValue {
	ret := make([]*domain.MoneyValue, 0, len(in))
	for _, v := range in {
		ret = append(ret, MapFromMoneyValue(v))
	}

	return ret
}

func MapFromLastPrices(in []*investapi.LastPrice) []*domain.LastPrice {
	ret := make([]*domain.LastPrice, 0, len(in))
	for _, v := range in {
		ret = append(ret, MapFromLastPrice(v))
	}

	return ret
}

func MapFromLastPrice(in *investapi.LastPrice) *domain.LastPrice {
	return &domain.LastPrice{
		Figi:  in.Figi,
		Price: MapFromQuotation(in.Price),
		Time:  in.Time.AsTime(),
	}
}

func MapFromAccounts(in []*investapi.Account) []*domain.Account {
	ret := make([]*domain.Account, 0, len(in))
	for _, v := range in {
		ret = append(ret, MapFromAccount(v))
	}
	return ret
}

func MapFromPositionsResponse(in *investapi.PositionsResponse) *domain.PositionsResponse {
	return &domain.PositionsResponse{
		Money:                   MapFromMoneyValues(in.Money),
		Blocked:                 MapFromMoneyValues(in.Blocked),
		Securities:              MapFromPositionsSecuritiesList(in.Securities),
		LimitsLoadingInProgress: in.LimitsLoadingInProgress,
	}
}

func MapFromPositionsSecuritiesList(in []*investapi.PositionsSecurities) []*domain.PositionsSecurities {
	ret := make([]*domain.PositionsSecurities, 0, len(in))
	for _, v := range in {
		ret = append(ret, MapFromPositionsSecurities(v))
	}
	return ret
}
func MapFromPositionsSecurities(in *investapi.PositionsSecurities) *domain.PositionsSecurities {
	return &domain.PositionsSecurities{
		Figi:    in.Figi,
		Blocked: in.Blocked,
		Balance: in.Balance,
	}
}

func MapFromAccount(in *investapi.Account) *domain.Account {
	return &domain.Account{
		Id:          in.Id,
		Type:        MapFromAccountType(in.Type),
		Name:        in.Name,
		Status:      MapFromAccountStatus(in.Status),
		OpenedDate:  in.OpenedDate.AsTime(),
		ClosedDate:  in.ClosedDate.AsTime(),
		AccessLevel: MapFromAccessLevel(in.AccessLevel),
	}
}

func MapFromOrderStates(in []*investapi.OrderState) []*domain.OrderState {
	ret := make([]*domain.OrderState, 0, len(in))
	for _, v := range in {
		ret = append(ret, MapFromOrderState(v))
	}
	return ret
}

func MapFromOrderState(in *investapi.OrderState) *domain.OrderState {
	return &domain.OrderState{
		OrderId:               in.OrderId,
		ExecutionReportStatus: MapFromOrderExecutionReportStatus(in.ExecutionReportStatus),
		LotsRequested:         in.LotsRequested,
		LotsExecuted:          in.LotsExecuted,
		InitialOrderPrice:     MapFromMoneyValue(in.InitialOrderPrice),
		ExecutedOrderPrice:    MapFromMoneyValue(in.ExecutedOrderPrice),
		TotalOrderAmount:      MapFromMoneyValue(in.TotalOrderAmount),
		AveragePositionPrice:  MapFromMoneyValue(in.AveragePositionPrice),
		InitialCommission:     MapFromMoneyValue(in.InitialCommission),
		ExecutedCommission:    MapFromMoneyValue(in.ExecutedCommission),
		Figi:                  in.Figi,
		Direction:             MapFromOrderDirection(in.Direction),
		InitialSecurityPrice:  MapFromMoneyValue(in.InitialSecurityPrice),
		Stages:                MapFromOrderStages(in.Stages),
		ServiceCommission:     MapFromMoneyValue(in.ServiceCommission),
		Currency:              in.Currency,
		OrderType:             MapFromOrderType(in.OrderType),
		OrderDate:             in.OrderDate.AsTime(),
	}
}

func MapFromOrderStages(in []*investapi.OrderStage) []*domain.OrderStage {
	ret := make([]*domain.OrderStage, 0, len(in))
	for _, v := range in {
		ret = append(ret, MapFromOrderStage(v))
	}
	return ret
}

func MapFromOrderStage(in *investapi.OrderStage) *domain.OrderStage {
	return &domain.OrderStage{
		Price:    MapFromMoneyValue(in.Price),
		Quantity: in.Quantity,
		TradeId:  in.TradeId,
	}
}

func MapFromAccountType(in investapi.AccountType) domain.AccountType {
	switch in {
	case investapi.AccountType_ACCOUNT_TYPE_TINKOFF:
		return domain.AccountTypeTinkoff
	case investapi.AccountType_ACCOUNT_TYPE_TINKOFF_IIS:
		return domain.AccountTypeTinkoffIis
	case investapi.AccountType_ACCOUNT_TYPE_INVEST_BOX:
		return domain.AccountTypeInvestBox
	}
	return domain.AccountTypeUnspecified
}

func MapFromAccountStatus(in investapi.AccountStatus) domain.AccountStatus {
	switch in {
	case investapi.AccountStatus_ACCOUNT_STATUS_NEW:
		return domain.AccountStatusNew
	case investapi.AccountStatus_ACCOUNT_STATUS_OPEN:
		return domain.AccountStatusOpen
	case investapi.AccountStatus_ACCOUNT_STATUS_CLOSED:
		return domain.AccountStatusClosed
	}

	return domain.AccountStatusUnspecified
}

func MapFromAccessLevel(in investapi.AccessLevel) domain.AccessLevel {
	switch in {
	case investapi.AccessLevel_ACCOUNT_ACCESS_LEVEL_FULL_ACCESS:
		return domain.AccessLevelFullAccess
	case investapi.AccessLevel_ACCOUNT_ACCESS_LEVEL_READ_ONLY:
		return domain.AccessLevelReadOnly
	case investapi.AccessLevel_ACCOUNT_ACCESS_LEVEL_NO_ACCESS:
		return domain.AccessLevelNoAccess
	}

	return domain.AccessLevelUnspecified
}
