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

func removeDuplicates2(nums []int) int {
	i := 0

	for j := 0; j < len(nums); j++ {
		if j == 0 || nums[j] != nums[j-1] {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}