package middle

func countServers(grid [][]int) int {
	m := len(grid)
	count := map[int]int{}
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				count[i]++
				count[m+j]++
			}
		}
	}
	res := 0
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				if count[i] > 1 || count[m+j] > 1 {
					res++
				}
			}
		}
	}
	return res
}
