package warmup

/*
 * Complete the 'jumpingOnClouds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY c as parameter.
 */

func jumpingOnClouds(c []int32) int32 {
	var jumps int32

	for pos := 0; pos < len(c); {
		if len(c) > pos+2 && c[pos+2] == 0 {
			pos = pos + 2
		} else {
			pos++
		}

		jumps++
	}

	return jumps - 1
}
