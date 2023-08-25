package middle

import (
	"math"
)

func ArrayNesting(nums []int) int {
	vis := make([]bool, len(nums))
	res := 0.0
	for i := 0; i < len(nums); i++ {
		count := 0.0
		for !vis[i] {
			count++
			vis[i] = true
			i = nums[i]
		}
		res = math.Max(res, count)
	}
	return int(res)
}

func ArrayNesting_Optimize(nums []int) int {
	res := 0.0
	n := len(nums)
	for i := 0; i < n; i++ {
		count := 0.0
		j := i
		for nums[j] < n {
			temp := nums[j]
			count++
			nums[j] = n
			j = temp
		}
		res = math.Max(res, count)
	}
	return int(res)
}
