package codesignal

func firstNotRepeatingChar(s string) string {
	count := map[string]int{}

	for _, c := range s {
		count[string(c)] += 1
	}

	for _, c := range s {
		if count[string(c)] == 1 {
			return string(c)
		}
	}

	return "_"
}
