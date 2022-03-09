package fundamentals

// reverseString Reverse a slice of string
func reverseString(s []string) []string {
	for i := 0; i < len(s)/2; i++ {
		other := len(s) - i - 1
		temp := s[i]
		s[i] = s[other]
		s[other] = temp
	}

	return s
}
