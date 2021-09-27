package problem283

func moveZeroes(nums []int) []int {

	cnt := 0

	i := 0
	for i < len(nums) {

		if nums[i] == 0 {

			nums = append(nums[:i], nums[i+1:]...)
			cnt++
			continue
		}

		i++
	}

	for cnt != 0 {
		nums = append(nums, 0)
		cnt--
	}

	return nums
}

func moveZeroes2(nums []int) []int {

	cnt := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[cnt] = nums[i]
			cnt++
		}
	}

	for i := cnt; i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}

func moveZeroes3(nums []int) []int {
	cnt := 0
	i := 0
	for i < len(nums) - cnt {
		if nums[i] == 0 {
			nums = append(append(nums[:i], nums[i+1:]...),0)
			cnt++
		} else {
			i++
		}
	}

	return nums
}