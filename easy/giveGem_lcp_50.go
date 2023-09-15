package easy

func giveGem(gem []int, operations [][]int) int {
	for _, op := range operations {
		v := gem[op[0]] / 2
		gem[op[0]] -= v
		gem[op[1]] += v
	}
	mi, ma := gem[0], gem[0]
	for _, v := range gem {
		if v < mi {
			mi = v
		}
		if v > ma {
			ma = v
		}
	}
	return ma - mi
}
