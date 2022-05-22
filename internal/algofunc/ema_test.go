package algofunc

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestCalcEma(t *testing.T) {
	a := int64(10)
	b := int32(900000000)

	v := decimal.NewFromInt(a).Add(decimal.New(int64(b), -9))
	t.Error(v.String())
}
