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
