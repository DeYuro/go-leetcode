package problem35

func searchInsert(nums []int, target int) int {
	if target > nums[len(nums)-1] {
		return len(nums)
	}

	for l,r := 0, len(nums)-1; l <= r; {
		if nums[l] == target {
			return l
		}

		if nums[r] == target {
			return r
		}

		if target < nums[l]  {
			return l
		}

		if target > nums[r] {
			return r+1
		}

		l++
		r--

		if l >= r {
			if target > nums[l] {
				return r+1
			} else {
				return l
			}
		}
	}

	return 0
}

// It is O(N) but leetcode accept it
func searchInsert2(nums []int, target int) int {
	for k,v := range nums {
		if v >= target {
			return k
		}
	}
	return len(nums)
}

func searchInsert3(nums []int, target int) int {
	l := 0
	r := len(nums)-1
	for l < r {
		m := (l+r)/2
		if nums[m] == target{
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	if nums[l] < target {
		return l + 1
	} else {
		return l
	}
}