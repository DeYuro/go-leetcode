package problem7

import "math"

func reverse(x int) int  {
	var r int

	for x != 0 {
		r = r * 10 + x % 10
		x = x / 10
	}
	if r > math.MaxInt32 || r < math.MaxInt32 * -1 {
		r = 0
	}
	return r
}
