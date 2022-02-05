package problem14

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {

	sort.Ints(nums)
	res := [][]int{}
	n := len(nums)

	for i:=0; i < n - 3;i++ {
		fmt.Printf("for i := 0; i < %d; i now %d \n", n-3, i)
		if i >0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i+1; j < n-2;j++ {
			fmt.Printf("for j := 0; j < %d; j now %d \n", n-2, j)

			if j > i+1 && nums[j]==nums[j-1] {
				continue
			}

			//fmt.Printf("target-nums[i]-nums[j]: %d := %d-%d-%d \n", target-nums[i]-nums[j], target, nums[i],nums[j])
			//s1 := target-nums[i]-nums[j]

			l,r := j+1, n-1

			for l < r {
				//fmt.Printf("s2 := nums[l] + nums[r]: %d := %d+%d \n", nums[l] + nums[r], nums[l] , nums[r])
				sum :=nums[i] + nums[j] + nums[l] + nums[r]

				if sum == target {
					res = append(res,[]int{nums[i],nums[j],nums[l],nums[r]})
					for l < r && nums[r-1] == nums[r] {
						r--
					}
					for l < r && nums[l+1] == nums[l] {
						l++
					}
					l++
					r--
				} else if sum > target {
					for l < r && nums[r-1] == nums[r] {
						r--
					}
					r--
				} else {
					for l < r && nums[l+1] == nums[l] {
						l++
					}
					l++
				}
			}
		}
	}

	return res
}