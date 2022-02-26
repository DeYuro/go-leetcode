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


func addBinary2(a string, b string) string {
	s := 0
	carry := 0
	res := ""
	la := len(a) - 1
	lb := len(b) - 1
	for la >= 0 || lb >= 0 || carry != 0{
		s = carry
		if la >= 0 {
			s += int(a[la] - '0')
			la --
		}
		if lb >= 0 {
			s += int(b[lb] - '0')
			lb --
		}
		carry = s / 2
		res = string(rune(s%2+'0')) + res
	}
	return res
}

func addBinary3(a string, b string) string {
	carry := '0'
	n := len(a)
	m := len(b)
	str := ""
	for i,j := n-1, m-1 ; i >= 0 || j >= 0 || carry != '0'; {
		ab := '0'
		bb := '0'
		if i >= 0 {
			ab = rune(a[i])
			i--
		}

		if j >= 0 {
			bb = rune(b[j])
			j--
		}
		res := '0'
		res, carry = add(ab,bb, carry)

		str = string(res) + str
	}

	return str
}