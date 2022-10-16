package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsIsomorphic(t *testing.T) {
	tests := map[string]struct {
		input    []string
		expected bool
	}{
		"test 1": {
			input:    []string{"egg", "add"},
			expected: true,
		},
		"test 2": {
			input:    []string{"foo", "bar"},
			expected: false,
		},
		"test 3": {
			input:    []string{"paper", "title"},
			expected: true,
		},
		"test 4": {
			input:    []string{"badc", "baba"},
			expected: false,
		},
		"test 5": {
			input:    []string{"baba", "abba"},
			expected: false,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, isIsomorphic(tt.input[0], tt.input[1]))
		})
	}
}
