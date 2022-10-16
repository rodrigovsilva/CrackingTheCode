package leetcode

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var res int

	for i := 0; i < len(s); i++ {
		visited := map[string]bool{}

		for j := i; j < len(s); j++ {
			vj := fmt.Sprintf("%c", s[j])
			if visited[vj] == true {
				break
			} else {
				if res <= j-i+1 {
					res = j - i + 1
				}
				visited[vj] = true
			}
		}

		vi := fmt.Sprintf("%c", s[i])

		visited[vi] = false

	}

	return res
}
