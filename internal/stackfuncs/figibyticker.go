package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type FigiByTicker struct {
}

func NewFigiByTicker() *FigiByTicker {
	return &FigiByTicker{}
}

func (f *FigiByTicker) Name() string {
	return "FigiByTicker"
}

func (f *FigiByTicker) BaseType() string {
	return "string"
}

func (f *FigiByTicker) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	return options.GetString("ticker")
}

func (f *FigiByTicker) Arguments() []*engine.Argument {
	return []*engine.Argument{
		{
			Id:           "ticker",
			Name:         "Ticker",
			Desc:         "Например, TCS",
			BaseType:     "string",
			ExtendedType: "figi-select",
			Required:     true,
			Value:        "",
		},
	}
}
