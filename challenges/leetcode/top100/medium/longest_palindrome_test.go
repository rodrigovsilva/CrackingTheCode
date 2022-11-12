package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"test 1": {
			input:    "babad",
			expected: "aba",
		},
		"test 2": {
			input:    "cbbd",
			expected: "bb",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, longestPalindrome(tt.input))
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected bool
	}{
		"test 1": {
			input:    "aba",
			expected: true,
		},
		"test 2": {
			input:    "bb",
			expected: true,
		},
		"test 3": {
			input:    "cbbd",
			expected: false,
		},
		"test 4": {
			input:    "c",
			expected: true,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.expected, isPalindrome(tt.input))
		})
	}
}
