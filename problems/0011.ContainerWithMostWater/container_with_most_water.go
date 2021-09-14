package problem11

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