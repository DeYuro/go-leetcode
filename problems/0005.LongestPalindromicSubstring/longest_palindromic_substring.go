package problem5

func longestPalindrome(s string) string {
	var answer string

	for i := 0; i < len(s); i++ {
		sub := s[i:]
		for j := len(sub); j > 0; j-- {
			if isPalindrome(sub[:j]) {
				if len(sub[:j]) > len(answer) {
					answer = sub[:j]
				}
			}
		}
	}

	return answer
}

func isPalindrome(s string) bool {
	for i,j := 0, len(s)-1; i < j; i,j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func longestPalindrome2(s string) string {
	longest := ""

	//babad
	for i := 0; i < len(s); i++ {
		// b
		// a->bab+
		// b->aba+->babad-
		odd := expandAndCount(s, i, i)
		// b
		// ab->baba-
		// ba->abad-
		even := expandAndCount(s, i, i+1)

		longer := ""
		if len(odd) > len(even) {
			longer = odd
		} else {
			longer = even
		}

		if len(longer) > len(longest) {
			longest = longer
		}
	}

	return longest
}

func expandAndCount(s string, start, end int) string {
	for start >= 0 && end <= len(s) - 1 && s[start] == s[end] {
		start -= 1
		end += 1
	}

	return s[start+1: end]
}

func longestPalindrome3(s string) string {
	if len(s) == 0 {
		return ""
	}

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	left, length := 0, 1
	for j := 0; j < len(s); j++ {
		dp[j][j] = true
		for i := 0; i < j; i++ {
			if j-i == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}
			if dp[i][j] && j-i+1 > length {
				length = j - i + 1
				left = i
			}
		}
	}
	return s[left : left+length]
}