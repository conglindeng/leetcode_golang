package difficult

import "math"

func minTrioDegree(n int, edges [][]int) int {
	g := make([][]int, n)
	degree := make([]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
	}
	for _, v := range edges {
		g[v[0]-1][v[1]-1] = 1
		g[v[1]-1][v[0]-1] = 1
		degree[v[0]-1]++
		degree[v[1]-1]++
	}
	res := math.MaxInt
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if g[i][j] != 1 {
				continue
			}
			for k := 0; k < n; k++ {
				if g[j][k] == 1 && g[k][i] == 1 {
					res = min(res, degree[i]+degree[j]+degree[k]-6)
				}
			}
		}
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}
