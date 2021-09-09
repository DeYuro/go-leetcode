package problem12

import "strings"

func intToRoman(i int) string {
	var roman []string

	counter := 1
	for i != 0 {
		r := i % 10
		i = i / 10

		if counter == 1 {
			roman = append(roman,convert("I", "V", "X", r))
		}

		if counter == 10 {
			roman = append(roman,convert("X", "L", "C", r))
		}

		if counter == 100 {
			roman = append(roman,convert("C", "D", "M", r))
		}

		if counter > 100 {
			roman = append(roman,convert("M", "M", "M", r))
			continue
		}
		counter *= 10
	}

	for i, j := 0, len(roman)-1; i < j; i, j = i+1, j-1 {
		roman[i], roman[j] = roman[j], roman[i]
	}

	return strings.Join(roman,"")
}

func convert(a,b,c string, num int) string {
	var roman []string
	if num == 9 {
		roman = append(roman,a+c)

	} else if num > 5 {
		roman = append(roman,b)
		for c := 0; c < num-5; c++ {
			roman = append(roman,a)
		}
	}else if num == 5 {
		roman = append(roman,b)
	}else if num == 4 {
		roman = append(roman,a+b)
	} else  {
		for c := 0; c < num; c++ {
			roman = append(roman,a)
		}
	}

	return strings.Join(roman,"")
}