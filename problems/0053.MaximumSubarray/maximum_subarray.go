package problem53

import "math"

// Kadane’s Algorithm
func maxSubArray(nums []int) int {
	max := math.MinInt64
	sumSeq := 0
	for _, v := range nums {
		sumSeq += v
		if max < sumSeq {
			max = sumSeq
		}

		if sumSeq < 0 {
			sumSeq = 0
		}
	}

	return max
}