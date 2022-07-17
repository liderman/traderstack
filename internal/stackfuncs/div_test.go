package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDiv_Run(t *testing.T) {
	fnc := NewDiv()

	tests := []struct {
		name   string
		a      decimal.Decimal
		b      decimal.Decimal
		result decimal.Decimal
	}{
		{"2/2=1", decimal.NewFromInt(2), decimal.NewFromInt(2), decimal.NewFromInt(1)},
		{"-2/2=-1", decimal.NewFromInt(-2), decimal.NewFromInt(2), decimal.NewFromInt(-1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := fnc.Arguments()
			args[0].Value = tt.a
			args[1].Value = tt.b

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
