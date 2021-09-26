package problem26

func removeDuplicates(nums []int) int {
	i := 0

	for i < len(nums) {
		if i == 0 {
			i++
			continue
		}

		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}

		i++
	}

	return len(nums)
}