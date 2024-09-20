package middle

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	end := points[0][1]
	count := 1
	for _, p := range points[1:] {
		if p[0] > end {
			end = p[1]
			count++
		}
	}
	return count
}
