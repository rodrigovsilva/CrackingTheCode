package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PivotIndex(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected int
	}{
		"test 1": {
			input:    []int{1, 7, 3, 6, 5, 6},
			expected: 3,
		},
		"test 2": {
			input:    []int{1, 2, 3},
			expected: -1,
		},
		"test 3": {
			input:    []int{2, 1, -1},
			expected: 0,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, pivotIndex(tt.input))
		})
	}
}
