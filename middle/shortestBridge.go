package middle

func ShortestBridge(grid [][]int) int {
	type pair struct {
		x, y int
	}
	directors := []pair{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	island := []pair{}
	q := []pair{}
	l := len(grid)
	step := 0
out:
	for i, row := range grid {
		for j, v := range row {
			if v == 0 {
				continue
			}
			q = append(q, pair{i, j})
			for len(q) > 0 {
				cur := q[0]
				island = append(island, cur)
				grid[cur.x][cur.y] = -1
				q = q[1:]
				for _, direct := range directors {
					m, n := cur.x+direct.x, cur.y+direct.y
					if m >= 0 && m < l && n < l && n >= 0 && grid[m][n] == 1 {
						q = append(q, pair{m, n})
					}
				}

			}
			break out
		}
	}
	q = island
	for {
		temp := q
		q = nil
		for len(temp) > 0 {
			cur := temp[0]
			temp = temp[1:]
			for _, direct := range directors {
				m, n := cur.x+direct.x, cur.y+direct.y
				if m < l && m >= 0 && n < l && n >= 0 {
					if grid[m][n] == 1 {
						return step
					} else if grid[m][n] == 0 {
						grid[m][n] = -1
						q = append(q, pair{m, n})
					}
				}
			}
		}
		step++
	}
}

func ShortestBridge_(grid [][]int) int {
	type pair struct {
		x, y int
	}
	directors := []pair{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	l := len(grid)
	step := 0
	for i, row := range grid {
		for j, v := range row {
			if v == 0 {
				continue
			}
			island := []pair{}
			q := []pair{{i, j}}
			grid[i][j] = -1
			for len(q) > 0 {
				cur := q[0]
				q = q[1:]
				island = append(island, cur)
				for _, direct := range directors {
					m, n := cur.x+direct.x, cur.y+direct.y
					if m >= 0 && m < l && n < l && n >= 0 && grid[m][n] == 1 {
						q = append(q, pair{m, n})
						grid[m][n] = -1
					}
				}
			}
			q = island
			for {
				temp := q
				q = nil
				for len(temp) > 0 {
					cur := temp[0]
					temp = temp[1:]
					for _, direct := range directors {
						m, n := cur.x+direct.x, cur.y+direct.y
						if m < l && m >= 0 && n < l && n >= 0 {
							if grid[m][n] == 1 {
								return step
							} else if grid[m][n] == 0 {
								grid[m][n] = -1
								q = append(q, pair{m, n})
							}
						}
					}
				}
				step++
			}
		}
	}
	return step
}
