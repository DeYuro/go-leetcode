package problem35

//
//func searchInsert(nums []int, target int) int {
//	if nums[len(nums)/2] >= target {
//		return helper(nums, len(nums)/2, len(nums) -1, target)
//	} else {
//		return helper(nums, 0, len(nums)/2, target)
//	}
//}
////[1,2,3,4,5,6,7,8,9,10]
//
//// t = 8
////[1,2,3,6,7,8,9]
//// [8,9]
//func helper(nums []int, l,r,t int) int {
//	if t == nums[l] {
//		return l
//	}
//
//	if t == nums[r] {
//		return r
//	}
//
//	if l == 0 {
//		return 0
//	}
//
//	if l >= len(nums)-1 {
//		return len(nums) - 1
//	}
//
//	if t > l
//	//if l == r {
//	//	if t > l {
//	//		return l + 1
//	//	} else
//	//}
//
//	return helper(nums, l, r, t)
//}

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