package middle

func maxValue(grid [][]int) int {
	//dp[i][j]=max{dp[i][j-1],dp[i-1][j]}+grid[i][j]
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, len(grid)+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j]) + grid[i][j]
		}
	}
	return dp[m][n]
}

func maxValue_(grid [][]int) int {
	//dp[i][j]=max{dp[i][j-1],dp[i-1][j]}+grid[i][j]
	m := len(grid)
	n := len(grid[0])
	dp := make([]int, len(grid[0])+1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[j+1] = max(dp[j+1], dp[j]) + grid[i][j]
		}
	}
	return dp[n]
}
