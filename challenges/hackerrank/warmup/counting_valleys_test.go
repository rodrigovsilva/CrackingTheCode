package warmup

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countingValleys(t *testing.T) {
	type input struct {
		steps int32
		path  string
	}

	tests := map[string]struct {
		input    input
		expected int32
	}{
		"test 1": {
			input: input{
				steps: 8,
				path:  "UDDDUDUU",
			},
			expected: 1,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, countingValleys(tt.input.steps, tt.input.path))
		})
	}
}
