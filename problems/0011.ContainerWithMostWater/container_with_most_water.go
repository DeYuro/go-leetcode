package problem11

// Work but "Time Limit Exceeded"
func maxArea(height []int) int {
	var area int

	max := func(a,b int) int {
		if a > b {
			return a
		}

		return b
	}

	for i := 0; i < len(height); i++ {
		for j := len(height) - 1; j > i; j-- {
			if height[i] <= height[j] {
				area = max(area,height[i] * (j - i))
			} else  {
				area = max(area,height[j] * (j - i))
			}
		}
	}

	return area
}


func maxArea2(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	maxF := func(a,b int) int {
		if a > b {
			return a
		}
		return b
	}

	minF := func(a,b int) int {
		if a < b {
			return a
		}
		return b
	}
	for left < right {
		max = maxF(max,(right-left) * minF(height[left], height[right]))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}