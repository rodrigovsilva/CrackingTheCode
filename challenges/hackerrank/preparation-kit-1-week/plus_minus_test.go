package preparation_kit_1_week

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusMinus(t *testing.T) {
	tests := map[string]struct {
		input            []int32
		expectedPositive string
		expectedNegative string
		expectedZero     string
	}{
		"main test": {
			input:            []int32{1, 1, 0, -1, -1},
			expectedPositive: "0.400000",
			expectedNegative: "0.400000",
			expectedZero:     "0.200000",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			p, n, z := plusMinus(tt.input)
			assert.Equal(t, tt.expectedPositive, p)
			assert.Equal(t, tt.expectedNegative, n)
			assert.Equal(t, tt.expectedZero, z)
		})

	}
}
