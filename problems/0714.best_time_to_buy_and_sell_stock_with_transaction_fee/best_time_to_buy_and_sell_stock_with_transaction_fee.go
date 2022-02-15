package problem714

func maxProfit(prices []int, fee int) int {
	bestBuy := make([]int, len(prices))
	bestSell := make([]int, len(prices))

	bestBuy[0] = prices[0]
	bestSell[0] = 0

	for i := 1 ; i < len(prices); i++ {
		bestBuy[i] = min(bestBuy[i-1], prices[i] - bestSell[i-1])
		bestSell[i] = max(bestSell[i-1], prices[i]-bestBuy[i-1]-fee )
	}

	return bestSell[len(prices)-1]
}

func max(a,b int) int  {
	if a > b {
		return a
	}

	return b
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}