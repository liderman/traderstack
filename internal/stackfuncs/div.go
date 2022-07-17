package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type Div struct {
}

func NewDiv() *Div {
	return &Div{}
}

func (l *Div) Name() string {
	return "/"
}

func (l *Div) BaseType() string {
	return engine.BaseTypeNumeric
}

func (l *Div) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	a, err := options.GetNumericDecimal("a")
	if err != nil {
		return nil, err
	}
	b, err := options.GetNumericDecimal("b")
	if err != nil {
		return nil, err
	}

	return a.Div(*b), nil
}

func (l *Div) Arguments() []*engine.Argument {
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
