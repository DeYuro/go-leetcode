package problem309

import "math"

func maxProfit(prices []int) int {
	sold := 0
	rest := 0
	hold := math.MinInt32
	for i :=0 ;i < len(prices); i++{
		prevSold := sold
		sold = hold + prices[i]

		rmp := rest - prices[i]

		hold = max(hold, rmp)
		rest = max(rest, prevSold)
	}
	return max(rest, sold)
}

func maxProfit3(prices []int) int {
	sold := 0
	hold := math.MinInt32
	for i:=0;i < len(prices);i++ {
		prevSold := sold
		sold = max(sold, prevSold + prices[i])
		hold = max(hold, prevSold - prices[i])
	}
	return sold
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	dp := make([]int, len(prices))
	ans := 0

	for i := 0; i < len(prices); i++ {
		maxProfit := 0
		for j := i - 1; j >= 0; j-- {
			maxProfit = max(maxProfit, prices[i]-prices[j])
			d := 0
			if j-2 >= 0 {
				d = dp[j-2]
			}
			dp[i] = max(dp[i], d+maxProfit)
		}

		ans = max(ans, dp[i])
	}

	return ans
}
