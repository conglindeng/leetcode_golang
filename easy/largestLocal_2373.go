package easy

func largestLocal(grid [][]int) [][]int {
	n := len(grid) - 3 + 1
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	getMax := func(m, n int) int {
		res := grid[m][n]
		for i := m; i < m+3; i++ {
			for j := n; j < n+3; j++ {
				res = max(res, grid[i][j])
			}
		}
		return res
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = getMax(i, j)
		}
	}
	return res
}
