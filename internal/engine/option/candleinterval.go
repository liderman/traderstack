package option

import "github.com/liderman/traderstack/internal/engine/baseoption"

func NewCandleInterval(id string) *baseoption.Integer {
	ret := baseoption.NewInteger(id)
	ret.ExtendedType = "candleInterval"
	ret.Name = "Интервал свечей"
	ret.Desc = "Например, 1 минута"
	return ret
}
