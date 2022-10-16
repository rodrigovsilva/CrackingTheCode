package leetcode

func runningSum(nums []int) []int {
	sums := make([]int, len(nums))

	sum := 0

	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		sums[i] = sum
	}

	return sums
}
