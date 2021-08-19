package main

import "fmt"

func main()  {
	fmt.Println(twoSumLoops([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
}

// Fast solution for example
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums)) // memory save

	for i, v := range nums {
		if idx, ok := m[target - v]; ok {
			return []int{idx, i}
		}
		m[v] = i
	}

	return nil // []int{}
}


// Slow solution for example
func twoSumLoops(nums []int, target int) []int {
	for i:= 0; i < len(nums); i++ {
		for j := i+1; j < len(nums);j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
