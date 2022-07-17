package stackfuncs

import (
	"errors"
	"fmt"
	"github.com/liderman/traderstack/internal/apiclient"
	"github.com/liderman/traderstack/internal/datamanager"
	"github.com/liderman/traderstack/internal/domain"
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type ActionStopLoss struct {
	api  apiclient.ApiClient
	md   *datamanager.MarketData
	inst *datamanager.Instruments
}

func NewActionStopLoss(
	api apiclient.ApiClient,
	md *datamanager.MarketData,
	inst *datamanager.Instruments,
) *ActionStopLoss {
	return &ActionStopLoss{
		api:  api,
		md:   md,
		inst: inst,
	}
}

func (a *ActionStopLoss) Name() string {
	return "ActionStopLoss"
}

func (a *ActionStopLoss) BaseType() string {
	return engine.BaseTypeBoolean
}

func (a *ActionStopLoss) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	figi, err := options.GetString("figi")
	if err != nil {
		return nil, err
	}
	percent, err := options.GetDecimal("percent")
	if err != nil {
		return nil, err
	}

	instrument, err := a.inst.GetShareByFigi(figi)
	if err != nil {
		return nil, err
	}

	err = a.checkWorkingExchange(instrument, now, isTest)
	if err != nil {
		return nil, err
	}

	inPortfolio, err := a.getPortfolioLots(instrument, accountId, isTest)
	if err != nil {
		return nil, err
	}

	if inPortfolio == nil {
		return false, nil
	}

	if !inPortfolio.ExpectedYield.IsNegative() {
		return false, nil
	}

	expectedYield := inPortfolio.ExpectedYield.Abs()

	if !expectedYield.GreaterThanOrEqual(percent) {
		return false, nil
	}

	fnc := a.api.PostOrder
	if isTest {
		fnc = a.api.PostSandboxOrder
	}
	resp, err := fnc(
		figi,
		inPortfolio.QuantityLots.IntPart(),
		inPortfolio.CurrentPrice.Value,
		domain.OrderDirectionSell,
		accountId,
		domain.OrderTypeMarket,
	)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания заказа: %w", err)
	}

	fmt.Printf("!!! StopLoss %s with %d lots and price %s\n", resp.Figi, resp.LotsRequested, inPortfolio.CurrentPrice.Value.String())
	return true, nil
}

func (a *ActionStopLoss) checkWorkingExchange(instrument *domain.Share, now time.Time, isTest bool) error {
	if isTest {
		to := now.Truncate(time.Minute).Add(time.Minute)
		from := now.Add(-time.Minute * 2)
		candles, err := a.md.GetCandles(instrument.Figi, from, to, domain.CandleInterval1Min)
		if err != nil {
			return err
		}

		if len(candles) > 0 {
			instrument.TradingStatus = domain.SecurityTradingStatusNormalTrading
		}
	}

	if !instrument.ApiTradeAvailableFlag {
		return errors.New("инструмент не доступен для торговли через API")
	}

	if !instrument.SellAvailableFlag {
		return errors.New("инструмент не доступен для продажи")
	}

	if !instrument.TradingStatus.AllowMarketOrder() {
		return errors.New("биржа закрыта для рыночных заявок")
	}

	return nil
}

func (a *ActionStopLoss) getPortfolioLots(instrument *domain.Share, accountId string, isTest bool) (*domain.PortfolioPosition, error) {
	fnc := a.api.GetPortfolio
	if isTest {
		fnc = a.api.GetSandboxPortfolio
	}
	portfolio, err := fnc(accountId)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения портфеля: %w", err)
	}

	for _, v := range portfolio {
		if v.Figi == instrument.Figi {
			return v, nil
		}
	}

	return nil, nil
}

func (a *ActionStopLoss) Arguments() []*engine.Argument {
	return []*engine.Argument{
		{
			Id:           "figi",
			Name:         "Figi",
			Desc:         "Например, TCS",
			BaseType:     "string",
			ExtendedType: "figi-select",
			Required:     true,
			Value:        "",
		},
		{
			Id:           "percent",
			Name:         "% убытка (decimal)",
			Desc:         "Например, 3",
			BaseType:     "decimal",
			ExtendedType: "",
			Required:     true,
			Value:        3,
		},
	}
}
