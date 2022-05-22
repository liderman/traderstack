package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type String struct {
}

func NewString() *String {
	return &String{}
}

func (s *String) Name() string {
	return "String"
}

func (s *String) BaseType() string {
	return "string"
}

func (s *String) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	return options.GetString("value")
}

func (s *String) Arguments() []*engine.Argument {
	return []*engine.Argument{
		{
			Id:           "value",
			Name:         "Строка",
			Desc:         "",
			BaseType:     "string",
			ExtendedType: "",
			Required:     true,
			Value:        "",
		},
	}
}
