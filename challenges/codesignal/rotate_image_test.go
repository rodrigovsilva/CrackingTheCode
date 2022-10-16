package codesignal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotateImage(t *testing.T) {
	assert.Equal(t, [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}, rotateImage([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
}
