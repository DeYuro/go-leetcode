package problem438

func findAnagrams(s string, p string) []int {
	res := []int{}
	ctr :=[26]int{}

	lenP := len(p)

	for _,c := range p {
		ctr[c - 'a']++
	}

	lenS := 1
	r := 0
	arr := [26]int{}
	for i:=r; i < len(s); {
		arr[s[i]-'a']++
		if lenS == lenP {
			lenS = 1

			if ctr == arr {
				res = append(res, r)
			}
			r++
			i = r
			arr = [26]int{}

		} else {
			i++
			lenS++
		}
	}

	return res
}

func findAnagrams2(s string, p string) []int {
	sArr, pArr := [26]int{}, [26]int{}
	ans := []int{}
	for _, v := range p {
		pArr[v - 'a']++
	}
	for i, v := range s {
		if i >= len(p) {
			sArr[s[i - len(p)] - 'a']--
		}
		sArr[v - 'a']++
		if sArr == pArr {
			ans = append(ans, i - len(p) + 1)
		}
	}
	return ans
}