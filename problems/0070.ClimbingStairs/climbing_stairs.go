package problem70


func climbStairsDp(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairsFib(n int) int {
	if n <= 2 {
		return n
	}

	a,b := 1, 2

	for i := 3; i <= n; i++ {
		save := b
		b = a + b
		a = save
	}

	return b
}