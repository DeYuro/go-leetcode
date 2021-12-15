package problem67

func addBinary(a string, b string) string {
	ra := Reverse(a)
	rb := Reverse(b)

	max := func(a,b int) int {
		if a > b {
			return a
		}

		return b
	}
	var result []rune
	maxLen := max(len(a),len(b))
	save := '0'

	for i := 0; i < maxLen; i++ {

		var av = '0'
		var bv = '0'
		if len(a) > i {
			av = rune(ra[i])
		} else {
		}
		if len(b) > i {
			bv = rune(rb[i])
		}

		v, s := add(av, bv, save)
		result = append([]rune{v}, result...)
		save = s
	}


	if save == '1' {
		result = append([]rune{'1'}, result...)
	}

	return string(result)
}


func add(a,b,c rune) (rune,rune)  {
	if a == '0' && b == '0' {
		return c, '0'
	}
	if a == '1' && b == '1' {
		return c, '1'
	}

	if c == '0' {
		return '1', '0'
	}

	return '0', '1'
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
