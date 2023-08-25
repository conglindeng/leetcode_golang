package easy

import "math"

func nearestValidPoint(x int, y int, points [][]int) int {
	dis := math.MaxInt
	index:=-1
	for m, i := range points {
		if x == i[0] || y == i[1] {
			if abs(x-i[0])+abs(y-i[1])<dis{
				index=m
				dis=abs(x-i[0])+abs(y-i[1])
			}
		}
	}
	return index
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
