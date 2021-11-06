package atm

import (
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func atm(amount int, bills map[int]int) []int {

	noms := sortedNoms(bills)
	test := collect(amount, noms, bills)
	_ = test
	return []int{}
}

func collect(amount int, noms []int, bills map[int]int) map[int]int {
	if amount == 0 {
		return map[int]int{}
	}

	if len(noms) == 0 {
		return map[int]int{}
	}
	curNom := noms[0]
	available := bills[curNom]
	need := amount / curNom
	number := min(available, need)
	for i := number; i >= 0; i-- {
		result := collect2(amount-i*curNom, noms[1:], bills)

		if len(result) > 0 {
			if i > 0 {
				m := map[int]int{curNom: i}
				for k, v := range result {
					m[k] = v
				}
				return m
			} else {
				return result
			}
		}
	}
	return nil
}


func sortedNoms(bills map[int]int) []int {
	var sorted []int

	for i := range bills {
		sorted = append(sorted, i)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))

	return sorted
}
