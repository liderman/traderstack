package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"time"
)

type Eq struct {
}

func NewEq() *Eq {
	return &Eq{}
}

func (l *Eq) Name() string {
	return "=="
}

func (l *Eq) BaseType() string {
	return engine.BaseTypeBoolean
}

func (l *Eq) Run(options *engine.Options, now time.Time, accountId string, isTest bool) (interface{}, error) {
	a, err := options.GetNumericDecimal("a")
	if err != nil {
		return nil, err
	}
	b, err := options.GetNumericDecimal("b")
	if err != nil {
		return nil, err
	}

	return a.Equal(*b), nil
}

func (l *Eq) Arguments() []*engine.Argument {
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
