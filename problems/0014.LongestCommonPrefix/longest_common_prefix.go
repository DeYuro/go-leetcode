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