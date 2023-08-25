package easy

import "math"

func minStartValue(nums []int) int {
	prefix := 0
	m := math.MaxInt
	for i := 0; i < len(nums); i++ {
		prefix += nums[i]
		if prefix < m {
			m = prefix
		}
	}
	if m > 0 {
		return 1
	}
	return -m + 1
}
