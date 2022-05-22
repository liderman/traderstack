package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"github.com/shopspring/decimal"
	"time"
)

type Decimal struct {
}

func NewDecimal() *Decimal {
	return &Decimal{}
}

func (b *Decimal) Name() string {
	return "Decimal"
}

func (b *Decimal) BaseType() string {
	return "decimal"
}

func (b *Decimal) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	return options.GetDecimal("value")
}

func (b *Decimal) Arguments() []*engine.Argument {
	return []*engine.Argument{
		{
			Id:           "value",
			Name:         "Число",
			Desc:         "",
			BaseType:     "decimal",
			ExtendedType: "",
			Required:     true,
			Value:        &decimal.Zero,
		},
	}
}
