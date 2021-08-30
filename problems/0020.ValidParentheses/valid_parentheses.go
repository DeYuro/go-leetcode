package problem20

func isValid(s string) bool {

	var stack []rune
	for _, v := range s {
		switch string(v) {
		case `(`:
			stack = append(stack, v)
		case `{`:
			stack = append(stack, v)
		case `[`:
			stack = append(stack, v)
		case `)`:
			if len(stack) > 0 && string(stack[len(stack)-1]) == `(` {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case `}`:
			if len(stack) > 0 && string(stack[len(stack)-1]) == `{` {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case `]`:
			if len(stack) > 0 && string(stack[len(stack)-1]) == `[` {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
