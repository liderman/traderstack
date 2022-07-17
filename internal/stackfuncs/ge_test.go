package stackfuncs

import (
	"github.com/liderman/traderstack/internal/engine"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGe_Run(t *testing.T) {
	ge := NewGe()

	tests := []struct {
		name   string
		a      int64
		b      int64
		result bool
	}{
		{"1 >= 2 = true", 2, 1, true},
		{"2 >= 2 = true", 2, 2, true},
		{"-2 >= -1 = true", -1, -2, true},
		{"2 >= 1 = false", 1, 2, false},
		{"-1 >= -2 = false", -2, -1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := ge.Arguments()
			args[0].Value = tt.a
			args[1].Value = tt.b

			opts := engine.NewOptions(args, map[string]interface{}{})
			resp, err := ge.Run(opts, time.Now(), "", false)

			assert.NoError(t, err)
			assert.Equal(t, resp, tt.result)
		})
	}

	opts := engine.NewOptions(ge.Arguments(), map[string]interface{}{})
	resp, err := ge.Run(opts, time.Now(), "", false)

	assert.Error(t, err)
	assert.Nil(t, resp)
}
