package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrays_rotLeft(t *testing.T) {
	type input struct {
		a []int32
		d int32
	}
	tests := map[string]struct {
		input    input
		expected []int32
	}{
		"main test": {
			input: input{
				a: []int32{1, 2, 3, 4, 5},
				d: 4,
			},
			expected: []int32{5, 1, 2, 3, 4},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, rotLeft(tt.input.a, tt.input.d))
		})

	}

}
