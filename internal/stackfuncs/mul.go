package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type Mul struct {
}

func NewMul() *Mul {
	return &Mul{}
}

func (l *Mul) Name() string {
	return "*"
}

func (l *Mul) BaseType() string {
	return engine.BaseTypeNumeric
}

func (l *Mul) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	a, err := options.GetNumericDecimal("a")
	if err != nil {
		return nil, err
	}
	b, err := options.GetNumericDecimal("b")
	if err != nil {
		return nil, err
	}

	return a.Mul(*b), nil
}

func (l *Mul) Arguments() []*engine.Argument {
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
