package middle

func checkValidGrid(grid [][]int) bool {
	if grid[0][0] != 0 {
		return false
	}
	n := len(grid[0]) * len(grid)
	position := make([][]int, n)
	for i, row := range grid {
		for j, v := range row {
			position[v] = []int{i, j}
		}
	}
	for i := 1; i < n; i++ {
		pre := position[i-1]
		cur := position[i]
		stepX := abs(pre[0] - cur[0])
		stepY := abs(pre[1] - cur[1])
		if stepX+stepY != 3 {
			return false
		}
		if stepX == 3 || stepY == 3 {
			return false
		}
	}
	return true
}
