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

type ActionTakeProfit struct {
	api  apiclient.ApiClient
	md   *datamanager.MarketData
	inst *datamanager.Instruments
}

func NewActionTakeProfit(
	api apiclient.ApiClient,
	md *datamanager.MarketData,
	inst *datamanager.Instruments,
) *ActionTakeProfit {
	return &ActionTakeProfit{
		api:  api,
		md:   md,
		inst: inst,
	}
}

func (a *ActionTakeProfit) Name() string {
	return "ActionTakeProfit"
}

func (a *ActionTakeProfit) BaseType() string {
	return "boolean"
}

func (a *ActionTakeProfit) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
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

	if !inPortfolio.ExpectedYield.GreaterThanOrEqual(*percent) {
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
		return nil, fmt.Errorf("ошибка создания заказа: %s", err)
	}

	fmt.Printf("!!! TakeProfit %s with %d lots and price %s\n", resp.Figi, resp.LotsRequested, inPortfolio.CurrentPrice.Value.String())
	return true, nil
}

func (a *ActionTakeProfit) checkWorkingExchange(instrument *domain.Share, now time.Time, isTest bool) error {
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

func (a *ActionTakeProfit) getPortfolioLots(instrument *domain.Share, accountId string, isTest bool) (*domain.PortfolioPosition, error) {
	fnc := a.api.GetPortfolio
	if isTest {
		fnc = a.api.GetSandboxPortfolio
	}
	portfolio, err := fnc(accountId)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения портфеля: %s", err)
	}

	for _, v := range portfolio {
		if v.Figi == instrument.Figi {
			return v, nil
		}
	}

	return nil, nil
}

func (a *ActionTakeProfit) Arguments() []*engine.Argument {
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
			Name:         "% роста (decimal)",
			Desc:         "Например, 15",
			BaseType:     "decimal",
			ExtendedType: "",
			Required:     true,
			Value:        15,
		},
	}
}
