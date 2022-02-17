package main

import (
	"fmt"
)

//Даны две строки строчных латинских символов: строка J и строка S. Символы, входящие в строку J, — «драгоценности», входящие в строку S — «камни». Нужно определить, какое количество символов из S одновременно являются «драгоценностями». Проще говоря, нужно проверить, какое количество символов из S входит в J.
// input
// ab
// aabbccd
//
// output
// 4
func main() {
	var s, j string

	fmt.Scanf("%s", &s)
	fmt.Scanf("%s", &j)

	result := 0

	gems := make(map[rune]bool)

	for _,v := range s {
		gems[v] = true
	}


	for _,v := range j {
		if _, ok := gems[v]; ok {
			result++
		}
	}

	fmt.Println(result)
}