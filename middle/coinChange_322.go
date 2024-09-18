package middle

import "sort"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = -1
	}
	dp[0] = 0
	sort.Ints(coins)
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if c > i {
				break
			}
			if dp[i-c] != -1 {
				if dp[i] == -1 {
					dp[i] = dp[i-c] + 1
				} else {
					dp[i] = min(dp[i], dp[i-c]+1)
				}
			}
		}
	}
	return dp[amount]
}
