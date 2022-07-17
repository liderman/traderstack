package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAbs_Run(t *testing.T) {
	fnc := NewAbs()

	tests := []struct {
		name   string
		value  decimal.Decimal
		result decimal.Decimal
	}{
		{"abs(9)=9", decimal.NewFromInt(9), decimal.NewFromInt(9)},
		{"abs(-9)=9", decimal.NewFromInt(-9), decimal.NewFromInt(9)},
		{"abs(0)=0", decimal.NewFromInt(0), decimal.NewFromInt(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := fnc.Arguments()
			args[0].Value = tt.value

			opts := engine.NewOptions(args, map[string]interface{}{})
			resp, err := fnc.Run(opts, time.Now(), "", false)

			assert.NoError(t, err)
			assert.Equal(t, tt.result.String(), resp.(decimal.Decimal).String())
		})
	}

	opts := engine.NewOptions(fnc.Arguments(), map[string]interface{}{})
	resp, err := fnc.Run(opts, time.Now(), "", false)

	assert.Error(t, err)
	assert.Nil(t, resp)
}
