package codesignal

func sudoku(grid [][]string) bool {

	var count int

	// rows
	for i := 0; i < len(grid); i++ {

		m := map[string]bool{}

		// iterate row
		for j := 0; j < len(grid); j++ {
			n := grid[i][j]

			if n == "." {
				continue
			}

			if m[n] {
				return false
			} else {
				count++
				m[n] = true
			}

		}
	}

	// cols
	for i := 0; i < len(grid); i++ {

		m := map[string]bool{}

		// iterate row
		for j := 0; j < len(grid); j++ {
			n := grid[j][i]

			if n == "." {
				continue
			}

			if m[n] {
				return false
			} else {
				count++
				m[n] = true
			}

		}

		if count == 9 {
			return true
		}
	}

	return true
}
