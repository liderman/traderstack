package stackfuncs

import (
	"fmt"
	"github.com/liderman/traderstack/internal/apiclient"
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type PortfolioLots struct {
	api apiclient.ApiClient
}

func NewPortfolioLots(api apiclient.ApiClient) *PortfolioLots {
	return &PortfolioLots{
		api: api,
	}
}

func (a *PortfolioLots) Name() string {
	return "PortfolioLots"
}

func (a *PortfolioLots) BaseType() string {
	return "integer"
}

func (a *PortfolioLots) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	figi, err := options.GetString("figi")
	if err != nil {
		return nil, err
	}

	fnc := a.api.GetPortfolio
	if isTest {
		fnc = a.api.GetSandboxPortfolio
	}
	portfolio, err := fnc(accountId)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения портфеля: %w", err)
	}

	for _, v := range portfolio {
		if v.Figi == figi {
			return v.QuantityLots.IntPart(), nil
		}
	}

	return int64(0), nil
}

func (a *PortfolioLots) Arguments() []*engine.Argument {
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
