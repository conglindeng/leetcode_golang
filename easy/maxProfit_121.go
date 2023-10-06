package easy

func maxProfit(prices []int) int {
	res := 0
	mi := prices[0]
	for i := 1; i < len(prices); i++ {
		res = max(res, prices[i]-mi)
		mi = min(mi, prices[i])
	}
	return res
}
