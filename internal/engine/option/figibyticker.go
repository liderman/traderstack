package option

import "github.com/liderman/traderstack/internal/engine/baseoption"

func NewFigiByTicker(id string) *baseoption.String {
	ret := baseoption.NewString(id)
	ret.ExtendedType = "figiByTicker"
	ret.Name = "Тикер"
	ret.Desc = "Например, TCSG"
	return ret
}
