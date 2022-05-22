package stackfuncs

import (
	"fmt"
	"github.com/liderman/traderstack/internal/apiclient"
	"github.com/liderman/traderstack/internal/domain"
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type InOrdersBuyMarketLots struct {
	api apiclient.ApiClient
}

func NewInOrdersBuyMarketLots(api apiclient.ApiClient) *InOrdersBuyMarketLots {
	return &InOrdersBuyMarketLots{
		api: api,
	}
}

func (a *InOrdersBuyMarketLots) Name() string {
	return "InOrdersBuyMarketLots"
}

func (a *InOrdersBuyMarketLots) BaseType() string {
	return "integer"
}

func (a *InOrdersBuyMarketLots) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	figi, err := options.GetString("figi")
	if err != nil {
		return nil, err
	}

	fnc := a.api.GetOrders
	if isTest {
		fnc = a.api.GetSandboxOrders
	}
	orders, err := fnc(accountId)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения списка заказов: %s", err)
	}

	lots := int64(0)
	for _, v := range orders {
		if v.OrderType == domain.OrderTypeMarket &&
			v.Direction == domain.OrderDirectionBuy &&
			v.ExecutionReportStatus.InProcess() &&
			v.Figi == figi {
			lots += v.LotsRequested
		}
	}

	return lots, nil
}

func (a *InOrdersBuyMarketLots) Arguments() []*engine.Argument {
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
	}
}
