package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type Sub struct {
}

func NewSub() *Sub {
	return &Sub{}
}

func (l *Sub) Name() string {
	return "-"
}

func (l *Sub) BaseType() string {
	return engine.BaseTypeNumeric
}

func (l *Sub) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	a, err := options.GetNumericDecimal("a")
	if err != nil {
		return nil, err
	}
	b, err := options.GetNumericDecimal("b")
	if err != nil {
		return nil, err
	}

	return a.Sub(*b), nil
}

func (l *Sub) Arguments() []*engine.Argument {
	return []*engine.Argument{
		{
			Id:           "a",
			Name:         "Слева",
			Desc:         "",
			BaseType:     "numeric",
			ExtendedType: "",
			Required:     true,
		},
		{
			Id:           "b",
			Name:         "Справа",
			Desc:         "",
			BaseType:     "numeric",
			ExtendedType: "",
			Required:     true,
		},
	}
}
