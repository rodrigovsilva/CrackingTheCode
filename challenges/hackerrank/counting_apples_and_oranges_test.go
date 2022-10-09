package hackerrank

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountingApplesAndOranges(t *testing.T) {

	type input struct {
		s       int32
		t       int32
		a       int32
		b       int32
		apples  []int32
		oranges []int32
	}

	type expected struct {
		applesCount  int
		orangesCount int
	}

	tests := map[string]struct {
		input    input
		expected expected
	}{
		"main test": {
			input: input{
				s:       7,
				t:       11,
				a:       5,
				b:       15,
				apples:  []int32{-2, 2, 1},
				oranges: []int32{5, -6},
			},
			expected: expected{
				applesCount:  1,
				orangesCount: 1,
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			applesCount, orangesCount := countApplesAndOranges(tt.input.s, tt.input.t, tt.input.a, tt.input.b, tt.input.apples, tt.input.oranges)
			assert.Equal(t, tt.expected.applesCount, applesCount)
			assert.Equal(t, tt.expected.orangesCount, orangesCount)
		})
	}

}
