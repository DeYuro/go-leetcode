package problem17

func letterCombinations(digits string) []string {


	if digits == "" {
		return []string{}
	}

	m := map[byte]string {
		'2':"abc",
		'3':"def",
		'4':"ghi",
		'5':"jkl",
		'6':"mno",
		'7':"pqrs",
		'8':"tuv",
		'9':"wxyz",
	}

	result := []string{""}

	for _ , d := range digits {

		str := m[byte(d)]
		var combination []string

		for _ , c := range str {

			char := c
			for _, v := range result {
				combination = append(combination, v+string(char))
			}

		}


		result = combination
	}


	return result
}
