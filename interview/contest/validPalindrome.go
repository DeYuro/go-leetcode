package main

import "unicode"

func isPalindrome(s string) bool {
	i, j := 0, len(s) - 1
	for i < j {
		if !isAlphaNumeric(s[i]) { i++; continue }
		if !isAlphaNumeric(s[j]) { j--; continue }
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) { return false }
		i, j = i + 1, j - 1
	}
	return true
}

func isAlphaNumeric(b byte) bool {
	return unicode.IsLetter(rune(b)) || unicode.IsNumber(rune(b))
}
