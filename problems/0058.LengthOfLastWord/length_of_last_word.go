package problem58

func lengthOfLastWord(s string) int {

	ss := []string{}

	tmp := ""
	for _, v := range s {
		if v == ' ' {
			if tmp != "" {
				ss = append(ss, tmp)
				tmp = ""
			}
		} else {
			tmp += string(v)
		}
	}

	if tmp != "" {
		ss = append(ss, tmp)
	}

	return len(ss[len(ss)-1])
}


func lengthOfLastWord2(s string) int {
	c := 0

	for i := len(s)-1; i >= 0; i-- {
		if s[i] != ' ' {
			c++
		} else if c > 0 {
			return c
		}
	}

	return c
}
