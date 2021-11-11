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

func atm(amount int, bills map[int]int) map[int]int {

	noms := sortedNoms(bills)
	test := collect(amount, noms, bills)
	return test
}

func collect(amount int, noms []int, bills map[int]int) map[int]int {
	if amount == 0 {
		return map[int]int{}
	}

	if len(noms) == 0 {
		return nil
	}
	curNom := noms[0]
	available := bills[curNom]
	need := amount / curNom
	number := min(available, need)
	for i := number; i >= 0; i-- {
		result := collect(amount-i*curNom, noms[1:], bills)

		if result != nil {
			if i > 0 {
				m := map[int]int{curNom: i}
				return mergeMap(m, result)
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

func mergeMap(m1, m2 map[int]int) map[int]int  {
	for k, v := range m2 {
		m1[k] = v
	}

	return m1
}