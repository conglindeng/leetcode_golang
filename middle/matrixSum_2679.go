package middle

import (
	"math"
	"sort"
)

func matrixSum(nums [][]int) int {
	for _, t := range nums {
		sort.Ints(t)
	}
	res := 0
	m, n := len(nums), len(nums[0])
	for i := 0; i < n; i++ {
		max := math.MinInt
		for j := 0; j < m; j++ {
			if nums[j][i] > max {
				max = nums[j][i]
			}
		}
		res += max
	}
	return res
}
