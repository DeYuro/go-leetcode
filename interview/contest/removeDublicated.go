package main

import "fmt"



func main() {
	var n int

	fmt.Scanf("%d", &n)

	nums := []int{}

	var num int

	for n > 0 {
		fmt.Scanf("%d", &num)
		nums = append(nums, num)
		n--
	}

	for i := 0; i < len(nums); i++{
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		fmt.Println(nums[i])
	}
}


//Дан упорядоченный по неубыванию массив целых 32-разрядных чисел. Требуется удалить из него все повторения.
//
//Желательно получить решение, которое не считывает входной файл целиком в память, т.е., использует лишь константный объем памяти в процессе работы.

//5
//2
//4
//8
//8
//8

//2
//4
//8