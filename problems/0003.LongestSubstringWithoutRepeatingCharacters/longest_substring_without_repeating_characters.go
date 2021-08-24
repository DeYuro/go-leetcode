package problem3

func lengthOfLongestSubstring(s string) int {

	m := make(map[rune]interface{})
	var max int
	var localMax int

	for _,v := range s {
		if _, ok := m[v]; !ok {
			localMax++
			m[v] = struct {}{}
		} else {
			if max < localMax {
				max = localMax
			}
			localMax = 1
			m = make(map[rune]interface{})
			m[v] = struct {}{}
		}
	}

	if max < localMax {
		max = localMax
	}

	return max
}
