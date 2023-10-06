package middle

func maxProfit_714(prices []int, fee int) int {
	have, sold := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		have, sold = max(have, sold-prices[i]), max(sold, have+prices[i]-fee)
	}
	return sold
}
