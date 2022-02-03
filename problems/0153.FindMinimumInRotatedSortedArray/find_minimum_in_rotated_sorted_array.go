package problem153

func findMin(nums []int) int {
	l,h := 0, len(nums) - 1


	for l < h {

		m := (l+h)/2
		if nums[m] > nums[h] {
			l = m + 1
		} else {
			h = m
		}
	}

	return nums[l]
}

func findMin2(nums []int) int {
	l, h :=0, len(nums)-1

	for l < h {
		// It's already sorted. so return the lowest from left.
		if nums[l]<= nums[h]{
			return nums[l]
		}
		m := (l + h)/2
		if nums[m] >= nums[l]{
			l = m + 1
		} else{
			h = m
		}
	}
	return nums[l]
}
//0,1,2,4,5,6,7
//1,2,4,5,6,7,0
//2,4,5,6,7,0,1