package problem15

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i+1
		r := n-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if sum == 0 {
				res = append(res,[]int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				r--
			}

			if sum > 0 {
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				r--
			}
			if sum < 0 {
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++

			}
		}
	}

	return res
}