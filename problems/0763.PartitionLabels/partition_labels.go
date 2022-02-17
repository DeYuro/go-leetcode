package problem763

func partitionLabels1(s string) []int {
	res := []int{}
	m := make(map[rune]int)

	for i, v := range s {
		m[v] = i
	}

	i := 0
	end := 0
	start := 0
	for i < len(s) {
		end = max(end, m[rune(s[i])])
		if i == end {
			res = append(res, i-start+1)
			start = i + 1
		}

		i++
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}


func partitionLabels2(s string) []int {

	res := []int{}

	m := make(map[rune]int)

	for i,v := range s {
		m[v] = i
	}

	count := 0
	end := 0
	for i,v := range s {
		count++

		if end < m[v] {
			end = m[v]
		}

		if i == end {
			res = append(res, count)
			count = 0
		}
	}

	return res
}