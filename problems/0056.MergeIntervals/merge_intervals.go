package problem56

import (
	"sort"
)

func merge(intervals [][]int) [][]int {

	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {return intervals[i][0] < intervals[j][0]})

	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]

		if last[1] >= intervals[i][0]  {
			if last[1] < intervals[i][1] {
				res[len(res) - 1] = []int{last[0], intervals[i][1]}
			}
		} else {
			res = append(res,intervals[i])
		}
	}

	return res
}