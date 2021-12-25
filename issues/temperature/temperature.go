package temperature

func temperatures(t []int)  []int {
	answer := make([]int, len(t))

	for i := 0; i < len(t); i++ {
		for j := i+1; j < len(t); j++ {
			if t[j] > t[i] {
				answer[i] = j-i
				break
			}
		}
	}
	return answer
}