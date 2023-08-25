package middle

import "math"

func minSideJumps(obstacles []int) int {
	const max = math.MaxInt / 2
	dp := [3]int{1, 0, 1}
	for _, i := range obstacles[1:] {
		m := max
		for j := 0; j < 3; j++ {
			if j == i-1 {
				dp[j] = max
			} else {
				m = min(m, dp[j])
			}
		}
		for j := 0; j < 3; j++ {
			if j != i-1 {
				dp[j] = min(dp[j], m+1)
			}
		}
	}
	return min(dp[0], min(dp[1], dp[2]))
}
