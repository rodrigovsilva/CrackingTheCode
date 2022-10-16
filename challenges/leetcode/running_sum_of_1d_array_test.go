package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RunningSums(t *testing.T) {

	tests := map[string]struct {
		input    []int
		expected []int
	}{
		"main test": {
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 3, 6, 10},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, runningSum(tt.input))
		})
	}
}
