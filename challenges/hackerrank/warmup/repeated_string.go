package warmup

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 * aba
 * abaabaabaa ba
 * 10
 */

func repeatedString(s string, n int64) int64 {
	loops := n / int64(len(s))
	rest := n - loops*int64(len(s))

	var totalPerWord int64

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "a" {
			totalPerWord++
		}
	}

	counter := loops * totalPerWord

	for i := 0; i < len(s[:rest]); i++ {
		if string(s[:rest][i]) == "a" {
			counter++
		}
	}

	return counter
}
