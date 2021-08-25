package problem3

func lengthOfLongestSubstring(s string) int {

	m := make(map[uint8]int)
	var max int
	var localMax int

	for i := 0; i < len(s); i++ {
		if idx, ok := m[s[i]]; !ok {
			localMax++
			m[s[i]] = i
		} else {
			if max < localMax {
				max = localMax
			}
			i = idx
			localMax = 0
			m = make(map[uint8]int)
		}
	}

	if max < localMax {
		max = localMax
	}

	return max
}

func lengthOfLongestSubstring2(s string) int {
	m, max, left := make(map[rune]int), 0, 0
	for idx, c := range s {
		if _, okay := m[c]; okay == true && m[c] >= left {
			if idx-left > max {
				max = idx - left
			}
			left = m[c] + 1
		}
		m[c] = idx
	}
	if len(s)-left > max {
		max = len(s) - left
	}
	return max
}