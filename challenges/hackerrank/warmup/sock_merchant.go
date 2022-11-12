package warmup

func sockMerchant(n int32, ar []int32) int32 {
	socksByColor := map[int32]int{}

	var pairs int32

	for i := 0; i < int(n); i++ {
		if socksByColor[ar[i]] == 0 {
			socksByColor[ar[i]] = 1
		} else {
			socksByColor[ar[i]] = socksByColor[ar[i]] + 1
		}

		if socksByColor[ar[i]]%2 == 0 {
			pairs++
		}
	}

	return pairs
}
