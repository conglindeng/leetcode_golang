package middle

import "sort"

func eliminateMaximum(dist []int, speed []int) int {
	//todo:dcl

	n := len(dist)
	tmp := make([]int, n)
	for i := 0; i < n; i++ {
		tmp[i] = (dist[i]-1)/speed[i] + 1
	}
	sort.Ints(tmp)
	for i := 0; i < n; i++ {
		if tmp[i] <= i {
			return i
		}
	}
	return n
}
