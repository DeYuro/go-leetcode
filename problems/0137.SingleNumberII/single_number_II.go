package problem136

func singleNumber(nums []int) int {
	m := make(map[int]int)


	for _,v := range nums {
		m[v]++
	}

	for i, v := range m{
		if v == 1 {
			return i
		}
	}

	return -1
}

func singleNumber2(nums []int) int {
	a := 0
	b := 0
	for _, v := range nums {
		a = a^v
		a = a^b
		b = b^v
		b = b^a
	}
	return a
}