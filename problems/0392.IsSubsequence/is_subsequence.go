package problem392

func isSubsequence(s string, t string) bool {
	i,j := 0,0

	for j < len(t) {

		if i == len(s) {
			return true
		}

		if t[j] == s[i] {
			i++
		}

		j++
	}

	return i == len(s)
}