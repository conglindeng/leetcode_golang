package middle

import "sort"

func MaxSumDivThree(nums []int) int {
	res := 0
	one := []int{}
	two := []int{}
	for _, n := range nums {
		if n%3 == 0 {
			res += n
		} else if n%3 == 1 {
			one = append(one, n)
		} else {
			two = append(two, n)
		}

	}
	sort.Ints(one)
	sort.Ints(two)
	


	return res

	//todo:dcl
}
