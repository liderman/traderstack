package algofunc

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcRsi(t *testing.T) {
	toDecimal := func(data []int64) []decimal.Decimal {
		var ret []decimal.Decimal
		for _, v := range data {
			ret = append(ret, decimal.NewFromInt(v))
		}
		return ret
	}

	tests := []struct {
		data     []int64
		expected string
	}{
		{
			data:     []int64{7, 6, 5, 4, 3, 2, 1, 0, 1, 2, 3, 4, 5, 6, 7},
			expected: "50",
		},
		{
			data:     []int64{0, 1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1, 0},
			expected: "50",
		},
		{
			data:     []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
			expected: "0",
		},
		{
			data:     []int64{14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: "100",
		},
		{
			data:     []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 10, 9, 8},
			expected: "23.0769230769230769",
		},
		{
			data:     []int64{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 2, 3, 4},
			expected: "76.9230769230769229",
		},
		{
			data:     []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expected: "100",
		},
		{
			data:     []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			expected: "100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			assert.Equal(t, tt.expected, CalcRsi(toDecimal(tt.data)).String())
		})
	}
}
