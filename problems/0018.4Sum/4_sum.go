package problem18

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {

	sort.Ints(nums)
	var res [][]int
	n := len(nums)

	for i:=0; i < n - 3;i++ {
		if i >0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i+1; j < n-2;j++ {

			if j > i+1 && nums[j]==nums[j-1] {
				continue
			}

			l,r := j+1, n-1

			for l < r {
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