package preparation_kit_1_week

import "fmt"

func plusMinus(arr []int32) (string, string, string) {
	var p, n, z int32

	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			p++
		} else if arr[i] < 0 {
			n++
		} else {
			z++
		}
	}

	size := len(arr)

	return fmt.Sprintf("%.6f", ratio(p, size)), fmt.Sprintf("%.6f", ratio(n, size)), fmt.Sprintf("%.6f", ratio(z, size))
}

func ratio(i int32, d int) float32 {
	v1 := float32(i)
	v2 := float32(d)

	return v1 / v2
}
