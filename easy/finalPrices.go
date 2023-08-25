package easy

func finalPrices(prices []int) []int {
	res := make([]int, len(prices))
	for index, num := range prices {
		if index < len(prices)-1 {
			m := 0
			for i := index + 1; i < len(prices); i++ {
				if prices[i] <= prices[index] {
					m = prices[i]
					break
				}
				res[index] = num - m
			}
		} else {
			res[index] = num
		}
	}
	return res
}

// minimal stack
func finalPrices_(prices []int) []int {
	l := len(prices)
	res := make([]int, l)
	stack := make([]int, 0)
	for i := l - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] > prices[i] {
			stack = stack[0 : len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = prices[i]
		} else {
			res[i] = prices[i] - stack[len(stack)-1]
		}
		stack = append(stack, prices[i])
	}
	return res
}
