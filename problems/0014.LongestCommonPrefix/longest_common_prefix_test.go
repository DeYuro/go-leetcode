package problem14

type test struct {
	input    []string
	expected string
}

func getTestcases() []test {
	return []test{
		{
		[]string{"flower","flow","flight"},
			"fl",
		},
		{
			[]string{"dog","racecar","car"},
			"",
		},
	}
}