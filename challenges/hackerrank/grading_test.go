package hackerrank

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrading(t *testing.T) {
	tests := map[string]struct {
		input    []int32
		expected []int32
	}{
		"first test": {
			input:    []int32{73, 67, 38, 33},
			expected: []int32{75, 67, 40, 33},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.expected, gradingStudents(tt.input))
		})
	}
}
