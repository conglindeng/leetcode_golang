package easy

func ShiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	count := m * n
	k %= count
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i, ints := range grid {
		for j := range ints {
			index := (i*n + j - k + count) % count
			res[i][j] = grid[index/n][index%n]
		}
	}
	return res
}