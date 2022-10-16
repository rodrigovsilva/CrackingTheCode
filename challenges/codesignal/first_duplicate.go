package codesignal

func solution1(a []int) int {
	numbers := map[int]int{}

	res := -1

	for i := 0; i < len(a); i++ {
		numbers[a[i]] = numbers[a[i]] + 1

		if numbers[a[i]] > 1 {
			res = a[i]

			break
		}
	}

	return res
}
