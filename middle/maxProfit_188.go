package middle

func maxProfit_188(k int, prices []int) int {
	tmp := make([]int, k*2)
	for i := 0; i < k*2; i += 2 {
		tmp[i] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		for j := len(tmp) - 1; j >= 0; j-- {
			if j%2 == 0 {
				if j != 0 {
					tmp[j] = max(tmp[j], tmp[j-1]-prices[i])
				} else {
					tmp[j] = max(tmp[j], -prices[i])
				}
			} else {
				tmp[j] = max(tmp[j], tmp[j-1]+prices[i])
			}
		}
	}
	return tmp[len(tmp)-1]
}
