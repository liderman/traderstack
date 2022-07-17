package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBoolean_Run(t *testing.T) {
	fnc := NewBoolean()
	var args []*engine.Argument

	opts := engine.NewOptions(args, map[string]interface{}{})
	result, err := fnc.Run(opts, time.Now(), "", false)

	assert.Error(t, err)
	assert.Equal(t, false, result)

	args = fnc.Arguments()
	args[0].Value = true
	opts = engine.NewOptions(args, map[string]interface{}{})
	result, err = fnc.Run(opts, time.Now(), "", false)

	assert.NoError(t, err)
	assert.Equal(t, true, result)
}
