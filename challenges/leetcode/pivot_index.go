package leetcode

func pivotIndex(nums []int) int {
	size := len(nums)

	totalSum := 0
	for i := 0; i < size; i++ {
		totalSum = totalSum + nums[i]
	}

	leftSum := 0
	for i := 0; i < size; i++ {
		leftSum = leftSum + nums[i]

		if leftSum == totalSum-leftSum+nums[i] {
			return i
		}
	}

	return -1
}
