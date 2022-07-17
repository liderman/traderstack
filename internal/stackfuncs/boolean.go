package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type Boolean struct {
}

func NewBoolean() *Boolean {
	return &Boolean{}
}

func (b *Boolean) Name() string {
	return "Boolean"
}

func (b *Boolean) BaseType() string {
	return engine.BaseTypeBoolean
}

func (b *Boolean) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	return options.GetBoolean("value")
}

func (b *Boolean) Arguments() []*engine.Argument {
	return []*engine.Argument{
		{
			Id:           "value",
			Name:         "Boolean",
			Desc:         "",
			BaseType:     "boolean",
			ExtendedType: "",
			Required:     true,
			Value:        false,
		},
	}
}
