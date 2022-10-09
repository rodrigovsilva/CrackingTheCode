package hackerrank

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtraLongFactorial(t *testing.T) {
	tests := map[string]struct {
		input    int32
		expected string
	}{
		"main test": {
			input:    25,
			expected: "15511210043330985984000000",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, extraLongFactorials(tt.input))
		})

	}

}