package middle

import "math"

func MaxSubarraySumCircular(nums []int) int {
	n := len(nums)
	res := math.MinInt
	for i := 0; i < n; i++ {
		cur := 0
		for j := 0; j < n; j++ {
			res = max(res, nums[(i+j)%n])
			cur += nums[(i+j)%n]
			if cur < 0 {
				cur = 0
			} else {
				res = max(res, cur)
			}
		}
	}
	return res
}

func MaxSubarraySumCircular_again(nums []int) int {
	// n := len(nums)
	// res := nums[]
	// sum := 0
	// index = 0
	
	//todo:dcl
	return -1
}
