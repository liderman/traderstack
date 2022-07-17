package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRound_Run(t *testing.T) {
	fnc := NewRound()

	tests := []struct {
		name   string
		value  decimal.Decimal
		places int64
		result decimal.Decimal
	}{
		{"round(5.45, 1)=5.5", decimal.NewFromFloat(5.45), 1, decimal.NewFromFloat(5.5)},
		{"round(545, -1)=550", decimal.NewFromInt(545), -1, decimal.NewFromInt(550)},
		{"round(555, 2)=555", decimal.NewFromInt(555), 2, decimal.NewFromInt(555)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := fnc.Arguments()
			args[0].Value = tt.value
			args[1].Value = tt.places

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
