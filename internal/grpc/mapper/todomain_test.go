package mapper

import (
	"github.com/liderman/traderstack/internal/grpc/gen/tinkoff/investapi"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MapFromQuotation(t *testing.T) {
	tests := []struct {
		actual   *investapi.Quotation
		expected string
	}{
		{
			actual: &investapi.Quotation{
				Units: 10,
				Nano:  900000000,
			},
			expected: "10.9",
		},
		{
			actual: &investapi.Quotation{
				Units: 74,
				Nano:  717500000,
			},
			expected: "74.7175",
		},
		{
			actual: &investapi.Quotation{
				Units: 73,
				Nano:  45000000,
			},
			expected: "73.045",
		},
		{
			actual: &investapi.Quotation{
				Units: -73,
				Nano:  -45000000,
			},
			expected: "-73.045",
		},
		{
			actual: &investapi.Quotation{
				Units: 0,
				Nano:  0,
			},
			expected: "0",
		},
		{
			actual:   nil,
			expected: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			assert.Equal(t, tt.expected, MapFromQuotation(tt.actual).String())
		})
	}
}
