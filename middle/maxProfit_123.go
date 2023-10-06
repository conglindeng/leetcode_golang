package middle

func maxProfit_123(prices []int) int {
	tmp := make([][4]int, len(prices))
	tmp[0][0] = -prices[0]
	tmp[0][2] = -prices[0]
	for i := 1; i < len(prices); i++ {
		tmp[i][0] = max(tmp[i-1][0], -prices[i])
		tmp[i][1] = max(tmp[i-1][1], tmp[i-1][0]+prices[i])
		tmp[i][2] = max(tmp[i-1][2], tmp[i-1][1]-prices[i])
		tmp[i][3] = max(tmp[i-1][3], tmp[i-1][2]+prices[i])
	}
	return tmp[len(prices)-1][3]
}

func maxProfit_123_space(prices []int) int {
	k := 2
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
