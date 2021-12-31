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

