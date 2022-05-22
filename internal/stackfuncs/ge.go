package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type Ge struct {
}

func NewGe() *Ge {
	return &Ge{}
}

func (g *Ge) Name() string {
	return ">="
}

func (g *Ge) BaseType() string {
	return "boolean"
}

func (g *Ge) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	a, err := options.GetNumericDecimal("a")
	if err != nil {
		return nil, err
	}
	b, err := options.GetNumericDecimal("b")
	if err != nil {
		return nil, err
	}

	return a.GreaterThanOrEqual(*b), nil
}

func (g *Ge) Arguments() []*engine.Argument {
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
