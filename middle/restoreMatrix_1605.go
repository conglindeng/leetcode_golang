package middle

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	n, m := len(rowSum), len(colSum)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}
	i, j := 0, 0
	for i < n && j < m {
		cur := min(rowSum[i], colSum[j])
		res[i][j] = cur
		rowSum[i] -= cur
		colSum[j] -= cur
		if rowSum[i] == 0 {
			i++
		}
		if colSum[j] == 0 {
			j++
		}
	}
	return res
}


