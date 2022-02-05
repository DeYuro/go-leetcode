package problem49

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	var res [][]string
	for _, v := range  strs {
		bs := []byte(v)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[i]
		})

		m[string(bs)] = append(m[string(bs)], v)
	}

	for _,v := range m {
		res = append(res, v)
	}

	return res
}

func groupAnagrams2(strs []string) [][]string {
	var result [][]string
	if len(strs) == 0 {
		return result
	}

	dict := make(map[string]int)
	for _, s := range strs {
		bs := []byte(s)

		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})

		sorted := string(bs)
		if idx, ok := dict[sorted]; ok {
			result[idx] = append(result[idx], s)
		} else {
			dict[sorted] = len(result)
			result = append(result, []string{s})
		}
	}

	return result
}

func groupAnagrams3(strs []string) [][]string {
	res := [][]string{}
	tmp := map[[26]int][]string{}
	for _, s := range strs {
		chars := [26]int{}
		for _, c := range s {
			chars[c - 'a']++
		}
		tmp[chars] = append(tmp[chars], s)
	}
	for _, v := range tmp {
		res = append(res, v)
	}
	return res
}

func groupAnagrams4(strs []string) [][]string {
	sorted := map[string][]int{}
	res := [][]string{}
	for i, v := range strs {
		tmp := []byte(v)
		sort.Slice(tmp, func(a, b int) bool { return tmp[a] < tmp[b] })
		sorted[string(tmp)] = append(sorted[string(tmp)], i)
	}
	for _, v := range sorted {
		tmp := []string{}
		for _, index := range v {
			tmp = append(tmp, strs[index])
		}
		res = append(res, tmp)
	}
	return res
}