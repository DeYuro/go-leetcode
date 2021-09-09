package problem13

func romanToInt(s string) int {
	var dec int
	roman := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		if i == 0 {
			dec =+ roman[s[i]]
			continue
		}
		//I can be placed before V (5) and X (10) to make 4 and 9.
		//X can be placed before L (50) and C (100) to make 40 and 90.
		//C can be placed before D (500) and M (1000) to make 400 and 900.
		if (s[i] == 'V' || s[i] == 'X') && s[i-1] == 'I' {
			dec+=roman[s[i]]-roman[s[i-1]]-1
			continue
		}
		if (s[i] == 'L' || s[i] == 'C') && s[i-1] == 'X' {
			dec+=roman[s[i]]-roman[s[i-1]]-10
			continue
		}

		if (s[i] == 'D' || s[i] == 'M') && s[i-1] == 'C' {
			dec+=roman[s[i]]-roman[s[i-1]]-100
			continue
		}
		dec+=roman[s[i]]
	}

	return dec
}
