package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDecimal_Run(t *testing.T) {
	fnc := NewDecimal()
	var args []*engine.Argument

	opts := engine.NewOptions(args, map[string]interface{}{})
	result, err := fnc.Run(opts, time.Now(), "", false)

	assert.Error(t, err)
	assert.Equal(t, decimal.Zero, result)

	args = fnc.Arguments()
	val := decimal.NewFromInt(100)
	args[0].Value = val
	opts = engine.NewOptions(args, map[string]interface{}{})
	result, err = fnc.Run(opts, time.Now(), "", false)

	assert.NoError(t, err)
	assert.Equal(t, val, result)
}
