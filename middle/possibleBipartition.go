package middle

func possibleBipartition(n int, dislikes [][]int) bool {
	dis := make([][]int, n)
	for _, num := range dislikes {
		x, y := num[0]-1, num[1]-1
		dis[x] = append(dis[x], y)
		dis[y] = append(dis[y], x)
	}
	color := make([]int, n)
	var color_dfs func(i, c int) bool
	color_dfs = func(i, c int) bool {
		color[i] = c
		for _, m := range dis[i] {
			if color[m] == c || (color[m] == 0 && !color_dfs(m, 3^c)) {
				return false
			}
		}
		return true
	}
	for _, i := range color {
		if color[i] == 0 && !color_dfs(i, 1) {
			return false
		}
	}
	return true
}

func possibleBipartition_bfs(n int, dislikes [][]int) bool {
	dis := make([][]int, n)
	for _, num := range dislikes {
		x, y := num[0]-1, num[1]-1
		dis[x] = append(dis[x], y)
		dis[y] = append(dis[y], x)
	}
	color := make([]int, n)
	for i, c := range color {
		if c == 0 {
			q := []int{i}
			color[i] = 1
			for len(q) > 0 {
				x := q[0]
				q = q[1:]
				for _, j := range dis[x] {
					if color[x] == color[j] {
						return false
					}
					if color[j] == 0 {
						color[j] = -color[x]
						q = append(q, j)
					}
				}
			}
		}
	}
	return true
}
