package atm

import (
	"sort"
)

func atm(amount int, bills map[int]int) []int {

	r := 1
	//save := amount
	var change []int
	var sorted []int

	for i := range bills {
		sorted = append(sorted, i)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	startPos := 0

	pos := 0

	for amount > 0 {
		i := pos
		for ;i < len(sorted); {
			if bills[sorted[i]] < 1 { //if  bills[i] not enough go to next bill
				i++
				continue
			}

			res := amount - sorted[i] // result  = amount - current bill value

			if res < 0 {
				// if result less then 0 or bills not enough go to next bill
				i++
				continue
			}

			amount = res // amount decrease
			bills[sorted[i]]--  // decrease bills
			pos = startPos // reset position

			change = append(change, sorted[i])
			continue
		}

		if pos < len(sorted) && len(change) > 0 {
			if amount == 0 {
				return change
			}

			for j := 1; j <= r; j++ {
				amount += change[len(change)-j]

			}


			//v := change[len(change)-1]
			//_ = v
			amount += change[len(change)-1]
			//amount = save

			//return bills to atm
			for _,v := range change {
				bills[v]++
			}

			bills[change[len(change)-1]]++
			change = []int{}
			//change = change[:len(change)-1]


			startPos = min(len(sorted)-1, startPos + 1)

			pos = startPos
		}

		if i > len(sorted) {
			return []int{}
		}
	}


	return change
}

func min(a,b int) int {
	if a < b {
		return a
	}

	return b
}

func atm2()  {
	
}

func collect(amount int, noms []int) []int {
	if amount == 0 {
		return []int{}
	}

	if len(noms) == 0 {
		return []int{}
	}

	curNom := noms[0]
	numsNotes := amount / curNom


}