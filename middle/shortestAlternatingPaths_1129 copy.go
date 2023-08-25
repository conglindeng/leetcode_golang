package middle

type edgeInfo struct {
	isRed bool
	point int
}

func ShortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	edges := map[int][]edgeInfo{}
	for _, edge := range redEdges {
		edges[edge[0]] = append(edges[edge[0]], edgeInfo{true, edge[1]})
	}
	for _, edge := range blueEdges {
		edges[edge[0]] = append(edges[edge[0]], edgeInfo{false, edge[1]})
	}
	res := make([]int, n)
	for i := 1; i < n; i++ {
		res[i] = -1
	}
	redAccessed := make([]bool, n)
	blueAccessed := make([]bool, n)
	var edgeDfs func(point, l int, isRed bool)
	edgeDfs = func(point, l int, isRed bool) {
		if accessed(redAccessed, blueAccessed, isRed, point) {
			return
		}
		if res[point] == -1 {
			res[point] = l
		} else {
			res[point] = min(res[point], l)
		}
		setAccessed(redAccessed, blueAccessed, isRed, point)
		for _, i := range edges[point] {
			if i.isRed == !isRed {
				edgeDfs(i.point, l+1, i.isRed)
			}
		}
		reSetAccessed(redAccessed, blueAccessed, isRed, point)
	}
	redAccessed[0] = true
	for _, i := range edges[0] {
		edgeDfs(i.point, 1, i.isRed)
	}
	return res
}

func accessed(redAccessed, blueAccessed []bool, isRed bool, point int) bool {
	if isRed {
		return redAccessed[point]
	} else {
		return blueAccessed[point]
	}
}

func setAccessed(redAccessed, blueAccessed []bool, isRed bool, point int) {
	if isRed {
		redAccessed[point] = true
	} else {
		blueAccessed[point] = true
	}
}

func reSetAccessed(redAccessed, blueAccessed []bool, isRed bool, point int) {
	if isRed {
		redAccessed[point] = false
	} else {
		blueAccessed[point] = false
	}
}
