package problem53

import "math"
// correct but Time Limit Exceeded
func maxSubArray(nums []int) int {
	max := math.MinInt64

	if len(nums) == 1 {
		return nums[0]
	}
	for i := 0; i < len(nums); i++ {
		localMax := nums[i]

		for j := i+1; j < len(nums); j++ {
			localMax += nums[j]
			if localMax > max {
				max = localMax
			}
		}
		if localMax > max {
			max = localMax
		}
	}

	return max
}