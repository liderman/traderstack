package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInteger_Run(t *testing.T) {
	fnc := NewInteger()
	var args []*engine.Argument

	opts := engine.NewOptions(args, map[string]interface{}{})
	result, err := fnc.Run(opts, time.Now(), "", false)

	assert.Error(t, err)
	assert.Equal(t, int64(0), result)

	args = fnc.Arguments()
	args[0].Value = int64(100)
	opts = engine.NewOptions(args, map[string]interface{}{})
	result, err = fnc.Run(opts, time.Now(), "", false)

	assert.NoError(t, err)
	assert.Equal(t, int64(100), result)
}
