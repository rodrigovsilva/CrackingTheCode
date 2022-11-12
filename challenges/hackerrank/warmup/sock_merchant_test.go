package warmup

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sockMerchant(t *testing.T) {
	type input struct {
		total int32
		items []int32
	}

	tests := map[string]struct {
		input    input
		expected int32
	}{
		"test 1": {
			input: input{
				total: 6,
				items: []int32{1, 2, 1, 2, 1, 3, 2},
			},
			expected: 2,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, sockMerchant(tt.input.total, tt.input.items))
		})
	}
}
