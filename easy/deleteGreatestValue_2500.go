package easy

import "sort"

func deleteGreatestValue(grid [][]int) int {
	for _, v := range grid {
		sort.Slice(v, func(i, j int) bool {
			return v[i] > v[j]
		})
	}
	n, m := len(grid[0]), len(grid)
	res := 0
	for i := 0; i < n; i++ {
		j := 0
		cur := grid[j][i]
		for ; j < m; j++ {
			cur = max(cur, grid[j][i])
		}
		res += cur
	}
	return res
}
