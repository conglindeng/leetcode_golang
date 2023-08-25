package middle

import "strconv"

func CountNicePairs(nums []int) int {
	const mod = 10e9 + 7
	res := 0
	for i, n := range nums {
		for j := i + 1; j < len(nums); j++ {
			if judge(n, nums[j]) {
				res++
				res = res % mod
			}
		}
	}
	return res
}

func judge(x, y int) bool {
	m := strconv.Itoa(x)
	n := strconv.Itoa(y)
	if len(m) != len(n) {
		return false
	}
	sum := int(m[0]) + int(n[len(n)-1])
	for i := 1; i < len(m); i++ {
		if int(m[i])+int(n[len(n)-1-i]) != sum {
			return false
		}
	}
	return true
}
