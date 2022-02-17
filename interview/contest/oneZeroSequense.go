package main

import "fmt"
//Требуется найти в бинарном векторе самую длинную последовательность единиц и вывести её длину.
//
//Желательно получить решение, работающее за линейное время и при этом проходящее по входному массиву только один раз.
func main() {
	var n int

	fmt.Scanf("%d", &n)

	nums := make([]int, n)

	var num int

	for n > 0 {
		fmt.Scanf("%d", &num)
		nums = append(nums, num)
		n--
	}

	m := 0
	c := 0
	for _, v := range nums {
		if v == 1 {
			c++
		} else {
			m = max(m,c)
			c = 0
		}
	}


	fmt.Println(max(m,c))
}


func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}