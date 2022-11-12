package warmup

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
	var v, seq int32

	for i := 0; i < int(steps); i++ {
		if string(path[i]) == "D" {
			seq--
			if seq == -1 {
				v++
			}
		}

		if string(path[i]) == "U" {
			seq++
		}
	}

	return v
}
