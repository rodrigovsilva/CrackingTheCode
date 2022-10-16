package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbers_1(t *testing.T) {
	expected := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 8,
			},
		},
	}

	assert.Equal(t, expected, addTwoNumbers(
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
		&ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 4,
				},
			},
		}))
}
