package problem154

import "sort"

func findMin(nums []int) int {
	l, r  := 0, len(nums) - 1

	for l <= r {
		m := (l+r) / 2

		if nums[l] < nums[r] || l == r {
			return nums[l]
		}

		if nums[m] == nums[l] {
			l++
			continue
		}

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func findMin2(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}