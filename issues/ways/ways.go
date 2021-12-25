package ways

func paths(n,m int) int {
	if n < 1 || m < 1 {
		return 0
	}

	if n == 1 && m == 1 {
		return 1
	}

	return paths(n-1, m) + paths(n, m-1)
}


func pathsDynamic(n,m int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	return helper(n,m, dp)
}

func helper(n,m int, dp [][]int) int {
	if n < 1 || m < 1 {
		return 0
	}

	if n == 1 && m == 1 {
		return 1
	}
	if dp[n][m] != 0{
		return dp[n][m]
	}

	dp[n][m] = helper(n-1, m, dp) + helper(n,m-1,dp)

	return dp[n][m]
}