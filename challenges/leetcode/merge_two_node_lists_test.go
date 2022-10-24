package leetcode

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {

	type input struct {
		n1 *ListNode
		n2 *ListNode
	}

	tests := map[string]struct {
		input    input
		expected *ListNode
	}{
		"test 1": {
			input{
				n1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
				n2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
								},
							},
						},
					},
				},
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			r := mergeTwoLists(tt.input.n1, tt.input.n2)

			j1, _ := json.Marshal(tt.expected)

			j2, _ := json.Marshal(r)

			fmt.Printf("%s\n%s", string(j1), string(j2))

			assert.Equal(t, tt.expected, r)
		})
	}

}
