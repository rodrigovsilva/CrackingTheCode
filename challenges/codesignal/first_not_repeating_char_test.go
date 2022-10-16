package codesignal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstNotRepeatingChar(t *testing.T) {
	assert.Equal(t, "c", firstNotRepeatingChar("abacabad"))
	assert.Equal(t, "_", firstNotRepeatingChar("abacabaabacaba"))
	assert.Equal(t, "d", firstNotRepeatingChar("abcdefghijklmnopqrstuvwxyziflskecznslkjfabe"))
}
