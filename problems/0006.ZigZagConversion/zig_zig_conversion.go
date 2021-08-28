package problem6

func convert(s string, numRows int) string {
	ss := make([][]rune, numRows)
	zig := false
	idx := 0
	if numRows == 1 {
		return s
	}

	for _, v := range s{
		ss[idx] = append(ss[idx], v)

		if idx + 1 < numRows && !zig {
			idx++
		}

		if zig && idx >= 0{
			idx--
		}

		if zig && idx == 0 {
			zig = false
		}

		if !zig && idx == numRows - 1 {
			zig = true
		}
	}

	var answer string
	for _, v := range ss {
		answer+=string(v)
	}

	return answer
}
//
// s = "PAYPALISHIRING", numRows = 3
//
// P     A     H     N
//  A  P  L  S  I  I   G
//   Y     I     R
//
// ANSWER = PAHNAPLSIIGYIR
//
// PERIOD = 2
//s = "PAYPALISHIRING", numRows = 4
//
// P       I      N
//  A    L  S    I  G
//   Y  A    H  R
//    P       I
//
// ANSWER PINALSIGYAHRPI
// PERIOD = 4

func convert2(s string, numRows int)string {
	if (numRows == 1) || (len(s) <= numRows) {
		return s
	}
	var period int = 2*numRows - 2
	byteStr := make([]byte, len(s))

	var i, j int
	j = 0
	for k := 0; k < numRows; k++ {
		i = k
		for i < len(s) {
			byteStr[j] = s[i]
			j++
			i += period - 2*(i%(numRows-1))
		}
	}
	return string(byteStr)
}