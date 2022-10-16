package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 2, lengthOfLongestSubstring("aab"))
	assert.Equal(t, 3, lengthOfLongestSubstring("abc"))
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	assert.Equal(t, 3, lengthOfLongestSubstring("dvdf"))

}
