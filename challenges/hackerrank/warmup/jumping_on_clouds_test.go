package warmup

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_jumpingOnClouds(t *testing.T) {
	type input struct {
		path []int32
	}

	tests := map[string]struct {
		input    input
		expected int32
	}{
		"test 1": {
			input: input{
				path: []int32{0, 0, 1, 0, 0, 1, 0},
			},
			expected: 4,
		},
		"test 2": {
			input: input{
				path: []int32{0, 0, 0, 1, 0, 0},
			},
			expected: 3,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, jumpingOnClouds(tt.input.path))
		})
	}
}
