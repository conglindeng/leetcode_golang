package middle

import "sort"

func minimumSize(nums []int, maxOperations int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return sort.Search(max, func(i int) bool {
		if i == 0 {
			return false
		}
		ops := 0
		for _, v := range nums {
			ops += (v - 1) / i
		}
		return ops <= maxOperations
	})
}
