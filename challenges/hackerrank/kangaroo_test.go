package hackerrank

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKangaroo(t *testing.T) {

	type input struct {
		k1 int32
		v1 int32
		k2 int32
		v2 int32
	}

	tests := map[string]struct {
		input    input
		expected string
	}{
		"main test": {
			input: input{
				k1: 0,
				v1: 3,
				k2: 4,
				v2: 2,
			},
			expected: "YES",
		},
		"main test 2": {
			input: input{
				k1: 0,
				v1: 2,
				k2: 5,
				v2: 3,
			},
			expected: "NO",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, kangaroo(tt.input.k1, tt.input.v1, tt.input.k2, tt.input.v2))
		})
	}

}
