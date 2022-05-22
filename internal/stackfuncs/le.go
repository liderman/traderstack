package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type Le struct {
}

func NewLe() *Le {
	return &Le{}
}

func (l *Le) Name() string {
	return "<="
}

func (l *Le) BaseType() string {
	return "boolean"
}

func (l *Le) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	a, err := options.GetNumericDecimal("a")
	if err != nil {
		return nil, err
	}
	b, err := options.GetNumericDecimal("b")
	if err != nil {
		return nil, err
	}

	return a.LessThanOrEqual(*b), nil
}

func (l *Le) Arguments() []*engine.Argument {
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
