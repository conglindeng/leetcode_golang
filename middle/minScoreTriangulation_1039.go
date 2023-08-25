package middle

import "math"

func minScoreTriangulation(values []int) int {
	//dp[i]
	cache := make(map[int]int)
	n := len(values)
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i+2 > j {
			return 0
		}
		if i+2 == j {
			return values[i] * values[i+1] * values[j]
		}
		key := i*n + j
		if _, ok := cache[key]; !ok {
			v := math.MaxInt
			for k := i + 1; k <= j; k++ {
				v = min(v, values[i]*values[j]*values[k]+dp(i, k)+dp(k, j))
			}
			cache[key] = v
		}
		return cache[key]
	}
	return dp(0, n-1)
}
