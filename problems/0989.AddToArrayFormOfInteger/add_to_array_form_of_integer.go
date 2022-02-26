package problem989

//inplace instead of
func addToArrayForm(num []int, k int) []int {
	carry := 0
	for i := len(num) - 1; i  >= 0; i-- {
		tmp := k % 10
		result := num[i] + tmp + carry
		num[i] = result % 10
		carry = result / 10
		k = k / 10
	}

	k = k + carry
	for k != 0 {
		num = append([]int{k%10}, num...)
		k /= 10
	}

	return num
}

func addToArrayForm2(num []int, k int) []int {
	res := []int{}

	for i := len(num) - 1; i >= 0; i-- {
		res = append(res, (num[i]+k)%10)
		k = (num[i] + k) / 10
	}
	for k > 0 {
		res = append(res, k%10)
		k /= 10
	}
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return res
}