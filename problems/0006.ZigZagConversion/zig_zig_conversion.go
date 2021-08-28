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
// P(0)     A(4)     	 H(8)     	N(12)
//  A(1)  P(3)  L(5)  S(7)  I(9)  I(11)  G(13)
//   Y(2)     		I(6)      R(10)
//
//   0+4+4+4
//   1+2+2+2+2+2+2
//   2+4+4
//
//  i := 0; next i=i+4-2*(0)? i%2 0%2=0  2=numRows-1
//  i := 1; next i=i+4-2*(1)? i%2 1%2=1
//  i := 2; next i=i+4-2*(0)? i%2 2%2=0
//

// ANSWER = PAHNAPLSIIGYIR
// PERIOD = 4 = 2*2(numRows-1)


//s = "PAYPALISHIRING", numRows = 4
//
// P(0)       I(6)      		N(12)
//  A(1)    L(5) S(7)    	I(11)  G(13)
//   Y(2)  A(4)    H(8)  R(10)
//      P(3)          I(9)
//
// 0+6+6
// 1+4+4+4+4
// 2+2+2+2
// 3+6
//  i := 0; next i=i+6-2*(0)? i%3 0%3=0   3=numRows-1
//  i := 1; next i=i+6-2*(1)? i%3 1%3=1
//  i := 2; next i=i+6-2*(2)? i%3 2%3=2
//  i := 3; next i=i+6-2*(0)? i%3 3%3=3
// ANSWER PINALSIGYAHRPI
// PERIOD = 6 = 2 * 3(numRows-1)

////s = "PAYPALISHIRING", numRows = 5
//
// P(0)				 H(8)
//	A(1)           S(7)	I(9)
//	  Y(2)		 I(6)	 R(10)
//      P(3)   L(5)			I(11)  G(13)
//         A(4)				  N(12)
// ANSWER PHASIYIRPLIGAN
// PERIOD = 8 = 2 * 4(numRows-1)
//
// 0+8(+8+8)
// 1+6+2(+6+2+6+2)
// 2+4+4(+4+4)
// 3+2+6+2(+6+2)
// 4+8(+8+8)

//  4 = numRows-1
//  i:=0; next i=i+8-2*(0)? i%4 0%4=0  i:=8; next i=i+8-2*(0) 8%4=0 ......
//  i:=1; next i=i+8-2*(1)? i%4 1%4=1  i:=7; next i=i+8-2*(3) 7%4=3;... next i=i+8-2*(1) 9%4=1
//  i:=2; next i=i+8-2*(2)? i%4 2%4=2  i:=6; next i=i+8-2*(2) 6%4=2 ......
//  i:=3; next i=i+8-2*(3)? i%4 3%4=3  i:=5; next i=i+8-2*(1) 5%4=1; i:=11; next i=i+8-2*(3) 11*4=3
//  i:=4; next i=i+8-2*(0)? i%4 4%4=0  i:=12; next i=i+8-2*(0) 12%4=0 .....

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