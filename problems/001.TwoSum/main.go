package main

import "fmt"

func main()  {
	fmt.Println(twoSumLoops([]int{2,7,11,15}, 9))
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
