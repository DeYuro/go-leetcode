package compress

import (
	"sort"
	"strconv"
)

func compress(nums []int) string {
	res := ""

	if len(nums) < 1 {
		return res
	}

	sort.Ints(nums)
	f := func(start, end int) {
		if len(res) > 0 {
			res += ","
		}
		if start == end {
			res += strconv.Itoa(end)
		} else  {
			res+= strconv.Itoa(start) + "-" + strconv.Itoa(end)
		}
	}

	start, end := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == end {
			end = nums[i]
			continue
		}
		f(start,end)

		end = nums[i]
		start = end
	}

	f(start,end)
	return res
}
