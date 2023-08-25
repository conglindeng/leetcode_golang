package middle

import "sort"

func getMaximumConsecutive(coins []int) int {
	res := 1
	sort.Ints(coins)
	for _, n := range coins {
		if n > res {
			break
		}
		res += n
	}
	return res
}

