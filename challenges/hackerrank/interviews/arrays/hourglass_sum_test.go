package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hourglassSum(t *testing.T) {
	tests := map[string]struct {
		input    [][]int32
		expected int32
	}{
		"main test": {
			input: [][]int32{
				{-9, -9, -9, 1, 1, 1},
				{0, -9, 0, 4, 3, 2},
				{-9, -9, -9, 1, 2, 3},
				{0, 0, 8, 6, 6, 0},
				{0, 0, 0, -2, 0, 0},
				{0, 0, 1, 2, 4, 0}},
			expected: 28,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, hourglassSum(tt.input))
		})

	}
}
