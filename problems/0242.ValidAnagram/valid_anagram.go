package problem0242

import (
	"reflect"
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	ss := []rune(s)
	st := []rune(t)


	sort.Slice(ss, func(i int, j int) bool { return ss[i] < ss[j] })
	sort.Slice(st, func(i int, j int) bool { return st[i] < st[j] })

	return reflect.DeepEqual(ss,st)
}


func isAnagram2(s string, t string) bool {
	ss := strings.Split(s, "")
	st := strings.Split(t, "")

	sort.Strings(ss)
	sort.Strings(st)

	return reflect.DeepEqual(ss, st)
}

// leetcode best solution symbol counter
func isAnagram3(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	table := [26]int{}
	for i := 0; i < len(s); i++ {
		table[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		table[t[i]-'a']--
		if table[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}