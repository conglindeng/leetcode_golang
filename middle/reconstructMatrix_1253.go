package middle

func ReconstructMatrix(upper int, lower int, colsum []int) [][]int {
	l := len(colsum)
	res := make([][]int, 2)
	res[0] = make([]int, l)
	res[1] = make([]int, l)
	sum := 0
	twoCount := 0
	for i := 0; i < l; i++ {
		n := colsum[i]
		if n < 0 || n > 2 {
			return [][]int{}
		}
		if n == 2 {
			res[0][i] = 1
			res[1][i] = 1
			twoCount++
		}
		sum += n
	}
	if upper+lower != sum {
		return [][]int{}
	}
	upper -= twoCount
	lower -= twoCount
	if lower < 0 || upper < 0 {
		return [][]int{}
	}
	for i, n := range colsum {
		if n == 1 {
			if upper > 0 {
				res[0][i] = 1
				upper--
			} else if lower > 0 {
				res[1][i] = 1
				lower--
			}
		}
	}
	return res
}
