package hackerrank

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) (int, int) {
	var applesCount, orangesCount int

	for i := 0; i < len(apples); i++ {
		if isLandedOnTheHouse(s, t, a+apples[i]) {
			applesCount++
		}
	}

	for i := 0; i < len(oranges); i++ {
		if isLandedOnTheHouse(s, t, b+oranges[i]) {
			orangesCount++
		}
	}

	///fmt.Println(applesCount)
	//fmt.Println(orangesCount)
	return applesCount, orangesCount
}

func isLandedOnTheHouse(s, t, d int32) bool {
	return s <= d && d <= t
}
