package middle

func maxProduct(words []string) int {
	res := 0
	masks := make([]int, len(words))
	for i, v := range words {
		val := 0
		for j := 0; j < len(v); j++ {
			val |= 1 << (v[j] - 'a')
		}
		masks[i] = val
	}
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			first := masks[i]
			second := masks[j]
			if first&second == 0 {
				res = max(res, len(words[i])*len(words[j]))
			}
		}
	}
	return res
}

func maxProduct_(words []string) int {
	masks := map[int]int{}
	for _, v := range words {
		val := 0
		for j := 0; j < len(v); j++ {
			val |= 1 << (v[j] - 'a')
		}
		if len(v) > masks[val] {
			masks[val] = len(v)
		}
	}
	res := 0
	for first, lF := range masks {
		for second, lS := range masks {
			if first&second == 0 {
				res = max(res, lS* lF)
			}
		}
	}
	return res
}
