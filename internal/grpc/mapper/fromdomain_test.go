package mapper

import (
	"github.com/liderman/traderstack/internal/grpc/gen/tinkoff/investapi"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapToQuotation(t *testing.T) {
	tests := []struct {
		actual   string
		expected *investapi.Quotation
	}{
		{
			actual: "10.9",
			expected: &investapi.Quotation{
				Units: 10,
				Nano:  900000000,
			},
		},
		{
			actual: "74.7175",
			expected: &investapi.Quotation{
				Units: 74,
				Nano:  717500000,
			},
		},
		{
			actual: "73.045",
			expected: &investapi.Quotation{
				Units: 73,
				Nano:  45000000,
			},
		},
		{
			actual: "-73.045",
			expected: &investapi.Quotation{
				Units: -73,
				Nano:  -45000000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.actual, func(t *testing.T) {
			val, _ := decimal.NewFromString(tt.actual)
			res := MapToQuotation(val)
			assert.Equal(t, tt.expected.Units, res.Units)
			assert.Equal(t, tt.expected.Nano, res.Nano)
		})
	}
}
