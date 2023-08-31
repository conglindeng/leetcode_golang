package middle

import "sort"

func NumFactoredBinaryTrees(arr []int) int {
	mod := int64(1e9 + 7)
	sort.Ints(arr)
	dp := make([]int64, len(arr))
	res := int64(0)
	for i, v := range arr {
		dp[i] = int64(1)
		left, right := 0, i-1
		for left <= right {
			if int64(arr[left])*int64(arr[right]) == int64(v) {
				if left == right {
					dp[i] = (dp[i] + dp[left]*dp[right]) % mod
				} else {
					dp[i] = (dp[i] + dp[left]*dp[right]*2) % mod
				}
				left++
				right--
			} else if int64(arr[left])*int64(arr[right]) < int64(v) {
				left++
			} else {
				right--
			}
		}
		res = (res + dp[i]) % mod
	}
	return int(res)
}
