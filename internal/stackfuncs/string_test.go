package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestString_Run(t *testing.T) {
	fnc := NewString()
	var args []*engine.Argument

	opts := engine.NewOptions(args, map[string]interface{}{})
	result, err := fnc.Run(opts, time.Now(), "", false)

	assert.Error(t, err)
	assert.Equal(t, "", result)

	args = fnc.Arguments()
	args[0].Value = "test"
	opts = engine.NewOptions(args, map[string]interface{}{})
	result, err = fnc.Run(opts, time.Now(), "", false)

	assert.NoError(t, err)
	assert.Equal(t, "test", result)
}
