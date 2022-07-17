package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type Mod struct {
}

func NewMod() *Mod {
	return &Mod{}
}

func (l *Mod) Name() string {
	return "mod"
}

func (l *Mod) BaseType() string {
	return engine.BaseTypeDecimal
}

func (l *Mod) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	a, err := options.GetNumericDecimal("a")
	if err != nil {
		return nil, err
	}
	b, err := options.GetNumericDecimal("b")
	if err != nil {
		return nil, err
	}

	return a.Mod(b), nil
}

func (l *Mod) Arguments() []*engine.Argument {
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
