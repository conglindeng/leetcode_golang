package middle

import "sort"

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	ans := 0
	n := len(reward1)
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		ans += reward2[i]
		diff[i] = reward1[i] - reward2[i]
	}
	sort.Ints(diff)
	for i := 1; i <= k; i++ {
		ans += diff[n-i]
	}
	return ans
}
