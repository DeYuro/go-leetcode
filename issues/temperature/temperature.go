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

type T struct {
	Val int
	Idx int
}

func temperatures2(t []int)  []int {
	answer := make([]int, len(t))

	stack := []T{}
	for i := len(t)-1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1].Val <= t[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			answer[i] = stack[len(stack)-1].Idx - i
		}

		stack = append(stack, T{Val: t[i], Idx: i})
	}
	return answer
}