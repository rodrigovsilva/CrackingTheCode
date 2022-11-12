package arrays

/*
 * Complete the 'hourglassSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func hourglassSum(arr [][]int32) int32 {
	h := len(arr)    //height
	w := len(arr[0]) // width

	var biggestSum *int32

	for i := 1; i < h-1; i++ {
		for j := 1; j < w-1; j++ {
			sum := arr[i-1][j-1] + arr[i-1][j] + arr[i-1][j+1]
			sum = sum + arr[i][j]
			sum = sum + arr[i+1][j-1] + arr[i+1][j] + arr[i+1][j+1]

			if biggestSum == nil {
				biggestSum = &sum
			} else if *biggestSum <= sum {
				biggestSum = &sum
			}
		}
	}
	return *biggestSum
}
