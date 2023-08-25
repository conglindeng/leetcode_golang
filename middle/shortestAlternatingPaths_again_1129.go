package middle

func ShortestAlternatingPaths_again(n int, redEdges [][]int, blueEdges [][]int) []int {
	type edgeInfo struct {
		x     int
		color int //0-red,1-blue
	}
	edges := make([][]edgeInfo, n)
	for _, edge := range redEdges {
		edges[edge[0]] = append(edges[edge[0]], edgeInfo{edge[1], 0})
	}
	for _, edge := range blueEdges {
		edges[edge[0]] = append(edges[edge[0]], edgeInfo{edge[1], 1})
	}
	res := make([]int, n)
	for i := 1; i < n; i++ {
		res[i] = -1
	}
	vis := make([][2]bool, n)
	vis[0] = [2]bool{true, true}
	work := []edgeInfo{{0, 0}, {0, 1}}
	for l := 0; len(work) > 0; l++ {
		temp := work
		work = nil
		for _, edge := range temp {
			x := edge.x
			if res[x] < 0 {
				res[x] = l
			}
			for _, next := range edges[x] {
				if edge.color != next.color && !vis[next.x][next.color] {
					vis[next.x][next.color] = true
					work = append(work, next)
				}
			}
		}
	}
	return res
}
