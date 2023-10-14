package middle

import "sort"

func sumDistance(nums []int, s string, d int) int {
	const mod int = 1e9 + 7
	n := len(nums)
	for i := 0; i < n; i++ {
		if s[i] == 'L' {
			nums[i] -= d
		} else {
			nums[i] += d
		}
	}
	sort.Ints(nums)
	res := 0
	for i := 1; i < n; i++ {
		res += (nums[i] - nums[i-1]) * i % mod * (n - i) % mod
		res = res % mod
	}
	return res
}
