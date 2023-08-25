package middle

func ShortestPathBinaryMatrix(grid [][]int) int {
	var dirs = [][]int{{0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {1, 0}, {0, 1}, {1, 1}, {1, -1}}
	m, n := len(grid), len(grid[0])
	if grid[0][0] != 0 {
		return -1
	}
	if m == 1 && n == 1 && grid[0][0] == 0 {
		return 1
	}
	arr := [][]int{{0, 0}}
	accessed := map[[2]int]bool{}
	accessed[[2]int{0, 0}] = true
	level := 1
	for len(arr) > 0 {
		tmp := [][]int{}
		for len(arr) > 0 {
			pop := arr[0]
			accessed[[2]int{pop[0], pop[1]}] = true
			arr = arr[1:]
			for _, dir := range dirs {
				x, y := pop[0]+dir[0], pop[1]+dir[1]
				if x < 0 || y < 0 || x >= m || y >= n {
					continue
				}
				if accessed[[2]int{x, y}] {
					continue
				}
				accessed[[2]int{x, y}] = true
				if grid[x][y] == 0 {
					if x == m-1 && y == n-1 {
						return level + 1
					} else {
						tmp = append(tmp, []int{x, y})
					}
				}
			}
		}
		level++
		arr = append(arr, tmp...)
	}
	return -1
}
