package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type Ne struct {
}

func NewNe() *Ne {
	return &Ne{}
}

func (l *Ne) Name() string {
	return "!="
}

func (l *Ne) BaseType() string {
	return engine.BaseTypeBoolean
}

func (l *Ne) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	a, err := options.GetNumericDecimal("a")
	if err != nil {
		return nil, err
	}
	b, err := options.GetNumericDecimal("b")
	if err != nil {
		return nil, err
	}

	return !a.Equal(*b), nil
}

func (l *Ne) Arguments() []*engine.Argument {
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
