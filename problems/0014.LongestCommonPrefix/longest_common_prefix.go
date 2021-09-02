package problem14

import "strings"

func longestCommonPrefix(strs []string) string {

	s := strs[0]
	longest := ""
	for i := 0; i < len(s); i++ {
		prefix := true
		ls := s[:i+1]
		for _,v := range strs[1:] {
			if strings.Index(v, ls) != 0 {
				prefix = false
			}
		}
		if prefix && len(longest) < len(ls) {
			longest = ls
		}
	}

	return longest
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := []byte{}
	baseLength, strLength := len(strs[0]), len(strs)

	for i := 0; i < baseLength; i++ {
		for j := 1; j < strLength; j++ {
			if len(strs[j]) <= i || strs[j][i] != strs[0][i]  {
				return string(prefix)
			}
		}

		prefix = append(prefix, strs[0][i])
	}

	return string(prefix)
}