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

type ActionSellMarket struct {
	api  apiclient.ApiClient
	md   *datamanager.MarketData
	inst *datamanager.Instruments
}

func NewActionSellMarket(
	api apiclient.ApiClient,
	md *datamanager.MarketData,
	inst *datamanager.Instruments,
) *ActionSellMarket {
	return &ActionSellMarket{
		api:  api,
		md:   md,
		inst: inst,
	}
}

func (a *ActionSellMarket) Name() string {
	return "ActionSellMarket"
}

func (a *ActionSellMarket) BaseType() string {
	return "boolean"
}

func (a *ActionSellMarket) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	condition, err := options.GetBoolean("condition")
	if err != nil {
		return nil, err
	}

	if !condition {
		return false, nil
	}
	figi, err := options.GetString("figi")
	if err != nil {
		return nil, err
	}
	lots, err := options.GetInteger("lots")
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

	lastPrice, err := a.md.GetLastPrice(figi)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения последней цены: %s", err)
	}

	fnc := a.api.PostOrder
	if isTest {
		fnc = a.api.PostSandboxOrder
	}
	resp, err := fnc(figi, lots, lastPrice.Price, domain.OrderDirectionSell, accountId, domain.OrderTypeMarket)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания заказа: %s", err)
	}

	fmt.Printf("!!! Sell %s with %d lots and price %s\n", resp.Figi, resp.LotsRequested, lastPrice.Price.String())
	return true, nil
}

func (a *ActionSellMarket) checkWorkingExchange(instrument *domain.Share, now time.Time, isTest bool) error {
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

func (a *ActionSellMarket) Arguments() []*engine.Argument {
	return []*engine.Argument{
		{
			Id:           "condition",
			Name:         "Условие",
			Desc:         "",
			BaseType:     "boolean",
			ExtendedType: "",
			Required:     true,
			Value:        false,
		},
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
			Id:           "lots",
			Name:         "Лотов",
			Desc:         "Например, 1",
			BaseType:     "integer",
			ExtendedType: "",
			Required:     true,
			Value:        1,
		},
	}
}
