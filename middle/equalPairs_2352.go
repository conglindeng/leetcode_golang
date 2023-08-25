package middle

import "fmt"

func equalPairs(grid [][]int) int {
	exist := map[string]int{}
	for _, row := range grid {
		exist[fmt.Sprint(row)]++
	}
	n := len(grid)
	res := 0
	for i := 0; i < n; i++ {
		var arr []int
		for j := 0; j < n; j++ {
			arr = append(arr, grid[j][i])
		}
		if val, ok := exist[fmt.Sprint(arr)]; ok {
			res += val
		}
	}
	return res
}
