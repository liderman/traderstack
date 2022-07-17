package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type Add struct {
}

func NewAdd() *Add {
	return &Add{}
}

func (l *Add) Name() string {
	return "+"
}

func (l *Add) BaseType() string {
	return engine.BaseTypeDecimal
}

func (l *Add) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	a, err := options.GetNumericDecimal("a")
	if err != nil {
		return nil, err
	}
	b, err := options.GetNumericDecimal("b")
	if err != nil {
		return nil, err
	}

	return a.Add(b), nil
}

func (l *Add) Arguments() []*engine.Argument {
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
