package medium

func longestPalindrome(s string) string {
	lStr := ""

	for i := 0; i < len(s); i++ {
		currStr := s[0:i]
		if isPalindrome(currStr) && len(lStr) < i+1 {
			lStr = currStr
		}
	}

	return lStr
}

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		if string(s[start]) != string(s[end]) {
			return false
		}

		start++
		end--
	}

	return true
}
