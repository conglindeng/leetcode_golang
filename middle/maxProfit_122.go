package middle

func maxProfit(prices []int) int {
	dp := make([][2]int, len(prices))
	dp[0][0], dp[0][1] = 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}

func maxProfit_dp_space(prices []int) int {
	sold, have := 0, -prices[0]
	for i := 0; i < len(prices); i++ {
		sold, have = max(sold, have+prices[i]), max(have, sold-prices[i])
	}
	return sold
}
