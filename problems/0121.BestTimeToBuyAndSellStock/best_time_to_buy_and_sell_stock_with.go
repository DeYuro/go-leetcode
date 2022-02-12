package problem121


func maxProfit(prices []int) int {
	profit := 0
	minPrice := prices[0]

	for i := 1;i < len(prices); i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		}

		if (prices[i] - minPrice) > profit {
			profit = prices[i] - minPrice
		}
	}

	return profit
}

