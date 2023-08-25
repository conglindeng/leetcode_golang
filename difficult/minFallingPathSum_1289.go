package difficult

import "math"

func minFallingPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = grid[i][j]
			} else {
				cur := math.MaxInt
				for k := 0; k < n; k++ {
					if j != k {
						cur = min(cur, dp[i-1][k])
					}
				}
				dp[i][j] = cur + grid[i][j]
			}
		}
	}
	res := math.MaxInt
	for j := 0; j < n; j++ {
		res = min(res, dp[m-1][j])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
