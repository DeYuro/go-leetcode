package problrem424

func characterReplacement(s string, k int) int {

	count := make(map[rune]int)

	res,l,r,m := 0,0,0,0

	for r < len(s) {
		count[rune(s[r])]++

		m = max(m,count[rune(s[r])])

		if r - l + 1 - m > k {
			count[rune(s[l])]--
			l++
		}

		res = max(res, r-l+1)
		r++
	}

	return res
}

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}

func characterReplacement2(s string, k int) int {

	count := [26]int{}

	res,l,r,m := 0,0,0,0

	for r < len(s) {
		count[s[r] - 'A']++

		m = max(m,count[s[r] - 'A'])

		if r - l + 1 - m > k {
			count[s[l] - 'A']--
			l++
		}

		res = max(res, r-l+1)
		r++
	}

	return res
}