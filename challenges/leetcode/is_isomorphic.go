package leetcode

func isIsomorphic(s string, t string) bool {
	mapS := map[string]string{}
	mapT := map[string]string{}

	for i := 0; i < len(s); i++ {
		sc := string(s[i])
		tc := string(t[i])

		sv, foundS := mapS[sc]

		tv, foundT := mapT[tc]

		if !foundS && !foundT {
			mapS[sc] = tc
			mapT[tc] = sc

			continue
		}

		if foundS && sv != tc {
			return false
		}

		if foundT && tv != sc {
			return false
		}

		if (!foundS && foundT) || (foundS && !foundT) {
			return false
		}
	}

	return true
}
