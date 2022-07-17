package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type Abs struct {
}

func NewAbs() *Abs {
	return &Abs{}
}

func (l *Abs) Name() string {
	return "abs"
}

func (l *Abs) BaseType() string {
	return engine.BaseTypeNumeric
}

func (l *Abs) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	value, err := options.GetNumericDecimal("value")
	if err != nil {
		return nil, err
	}

	return value.Abs(), nil
}

func (l *Abs) Arguments() []*engine.Argument {
	return []*engine.Argument{
		{
			Id:           "value",
			Name:         "Число",
			Desc:         "",
			BaseType:     "numeric",
			ExtendedType: "",
			Required:     true,
		},
	}
}
