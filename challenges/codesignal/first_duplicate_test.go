package codesignal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1_solution(t *testing.T) {
	assert.Equal(t, 3, solution1([]int{2, 1, 3, 5, 3, 2}))
	assert.Equal(t, 2, solution1([]int{2, 2}))
	assert.Equal(t, -1, solution1([]int{2, 4, 3, 5, 1}))
}
