package compressString

import "strconv"

func compressString(s string) string  {
	return string(compress([]byte(s)))
}

func compress(chars []byte)  []byte {
	var answer []byte
	var l byte

	f := func(ch byte, cn int) {
		answer = append(answer,ch)

		if cn == 1 {
			return
		}

		for _,v := range []byte(strconv.Itoa(cn)) {
			answer = append(answer, v)
		}
	}

	c := 1
	for _, v := range chars {
		if l == 0 {
			l = v
			continue
		}

		if v == l {
			c++
			continue
		}

		f(l,c)
		l = v
		c = 1
	}

	f(l,c)

	return answer
}
