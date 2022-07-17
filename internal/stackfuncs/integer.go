package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type Integer struct {
}

func NewInteger() *Integer {
	return &Integer{}
}

func (i *Integer) Name() string {
	return "Integer"
}

func (i *Integer) BaseType() string {
	return engine.BaseTypeInteger
}

func (i *Integer) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	return options.GetInteger("value")
}

func (i *Integer) Arguments() []*engine.Argument {
	return []*engine.Argument{
		{
			Id:           "value",
			Name:         "Число",
			Desc:         "",
			BaseType:     "integer",
			ExtendedType: "",
			Required:     true,
			Value:        0,
		},
	}
}
