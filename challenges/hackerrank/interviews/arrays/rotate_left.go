package arrays

func rotLeft(a []int32, d int32) []int32 {
	r := make([]int32, len(a))

	r = append(a[d:], a[:d]...)

	return r
}
