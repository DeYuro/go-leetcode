package problem136

func singleNumber(nums []int) int {
	m := make(map[int]interface{})

	for _, v := range nums {
		_, ok := m[v]

		if ok {
			delete(m,v)
			continue
		}


		m[v] = nil

	}

	for i := range m {
		return i
	}

	return 0
}


// xor solution
func singleNumber2(nums []int) int {
	var res int
	for _, num := range nums {
		res = res^num  // or res ^= num (xor)
		// example
		// 4 = 0100
		// 1 = 0001
		// 2 = 0010
		// 4,1,2,1,2
		// 0100 xor 0001 = 0101 (5)
		// 0101 xor 0010 = 0111 (7)
		// 0111 xor 0001 = 0110 (6)
		// 0110 xor 0010 = 0100 (4) = answer!
	}

	return res
}