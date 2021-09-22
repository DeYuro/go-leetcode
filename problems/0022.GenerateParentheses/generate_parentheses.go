package problem22

func generateParenthesis(n int) []string {
	if n == 0 { return []string{} }
	combinations := []string{}
	getCombination(n, n, "", &combinations)
	return combinations
}

func getCombination(left, right int, combo string, combinations *[]string) {

	if right < left {
		return //not valid combination
	}
	if left == 0 && right == 0 {
		*combinations = append(*combinations, combo)
	}

	// go down to zero and concat parentheses
	//
	// n = 2
	// left=2               right=2
	//     =1(
	//     =0(                           ---|
	//                      )    =1         |
	//                      )    =0         |
	//---------valid------------------      |
	//     =1(
	//                     )    =1
	//     =0(
	//                     )    =0
	//--------valid------------------
	//

	if left > 0 {
		getCombination(left-1, right, combo + "(", combinations)
	}
	if right > 0 {
		getCombination(left, right-1, combo + ")", combinations)
	}
}