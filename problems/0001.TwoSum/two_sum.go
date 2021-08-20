package problem1

// O(n)   time
// O(n)   memory
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums)) // memory save

	for i, v := range nums {
		if idx, ok := m[target - v]; ok {
			return []int{idx, i}
		}
		m[v] = i
	}

	return nil // []int{}
}


// O(n^2) time
// O(1)   memory
func twoSumLoops(nums []int, target int) []int {
	for i:= 0; i < len(nums); i++ {
		for j := i+1; j < len(nums);j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil // []int{}
}
