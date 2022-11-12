package warmup

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_repeatedString(t *testing.T) {
	type input struct {
		s string
		n int64
	}

	tests := map[string]struct {
		input    input
		expected int64
	}{
		"test 1": {
			input: input{
				s: "aba",
				n: 10,
			},
			expected: 7,
		},
		"test 2": {
			input: input{
				s: "a",
				n: 1000000000000,
			},
			expected: 1000000000000,
		},
		"test 3": {
			input: input{
				s: "aaaaaaaaa",
				n: 1000000000000,
			},
			expected: 1000000000000,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, repeatedString(tt.input.s, tt.input.n))
		})
	}
}
