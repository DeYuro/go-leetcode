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