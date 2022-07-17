package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type Round struct {
}

func NewRound() *Round {
	return &Round{}
}

func (l *Round) Name() string {
	return "round"
}

func (l *Round) BaseType() string {
	return engine.BaseTypeDecimal
}

func (l *Round) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	value, err := options.GetNumericDecimal("value")
	if err != nil {
		return nil, err
	}
	places, err := options.GetInteger("places")
	if err != nil {
		return nil, err
	}

	return value.Round(int32(places)), nil
}

func (l *Round) Arguments() []*engine.Argument {
	return []*engine.Argument{
		{
			Id:           "value",
			Name:         "Значение",
			Desc:         "",
			BaseType:     "numeric",
			ExtendedType: "",
			Required:     true,
		},
		{
			Id:           "places",
			Name:         "Округл.",
			Desc:         "",
			BaseType:     "integer",
			ExtendedType: "",
			Required:     true,
		},
	}
}
