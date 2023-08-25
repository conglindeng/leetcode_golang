package middle

import "math"

func PartitionDisjoint(nums []int) int {
	n := len(nums)
	leftMax := make([]int, n)
	rightMin := make([]int, n)
	ma := math.MinInt
	for i, n := range nums {
		if n > ma {
			ma = n
		}
		leftMax[i] = ma
	}
	mi := math.MaxInt
	for i := n - 1; i >= 0; i-- {
		if nums[i] < mi {
			mi = nums[i]
		}
		rightMin[i] = mi
	}
	for i := 0; i < n-1; i++ {
		if leftMax[i] <= rightMin[i+1] {
			return i + 1
		}
	}
	return n - 1
}