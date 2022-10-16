package codesignal

import "fmt"

func rotateImage(a [][]int) [][]int {
	n := len(a)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			temp := a[i][j]
			a[i][j] = a[j][i]
			a[j][i] = temp
		}

		low := 0
		high := n - 1
		for low < high {
			t := a[i][low]
			a[i][low] = a[i][high]
			a[i][high] = t
			low++
			high--
		}
	}
	/*
		for i := 0; i < n; i++ {
			low := 0
			high := n - 1
			for low < high {
				t := a[i][low]
				a[i][low] = a[i][high]
				a[i][high] = t
				low++
				high--
			}
		}*/

	fmt.Println(a)
	return a
}

/*
a =
[1, 2, 3],
[4, 5, 6],
[7, 8, 9]

s =
{7, 4, 1},
{8, 5, 2},
{9, 6, 3},
*/
