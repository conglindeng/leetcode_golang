package middle

import "strconv"

func QueensAttacktheKing(queens [][]int, king []int) [][]int {
	x, y := king[0], king[1]
	tmp := map[string][]int{}
	getFlag := func(n int) int {
		if n == 0 {
			return 0
		}
		if n > 0 {
			return 1
		}
		return -1
	}
	buildKey := func(x, y int) string {
		return strconv.Itoa(getFlag(x)) + "_" + strconv.Itoa(getFlag(y))
	}
	for _, v := range queens {
		qX, qY := v[0]-x, v[1]-y
		if qX == 0 || qY == 0 || abs(qX) == abs(qY) {
			k := buildKey(qX, qY)
			if v, ok := tmp[k]; ok {
				if v[0]*v[0]+v[1]*v[1] > qX*qX+qY*qY {
					tmp[k] = []int{qX, qY}
				}
			} else {
				tmp[k] = []int{qX, qY}
			}
		}
	}
	res := [][]int{}
	for _, v := range tmp {
		res = append(res, []int{v[0] + x, v[1] + y})
	}
	return res
}
