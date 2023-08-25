package middle

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[0][i] = matrix[0][i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			v := dp[i-1][j]
			if j-1 >= 0 {
				v = min(dp[i-1][j-1], v)
			}
			if j+1 < n {
				v = min(dp[i-1][j+1], v)
			}
			dp[i][j] = v + matrix[i][j]
		}
	}
	res := dp[n-1][0]
	for i := 0; i < n; i++ {
		res = min(res, dp[n-1][i])
	}
	return res
}
