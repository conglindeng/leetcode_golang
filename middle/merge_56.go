package middle

import "sort"

func merge_56(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	for _, v := range intervals {
		if len(res) == 0 {
			res = append(res, v)
		} else {
			if res[len(res)-1][0] <= v[0] && res[len(res)-1][1] >= v[0] {
				res[len(res)-1][1] = max(res[len(res)-1][1], v[1])
			} else {
				res = append(res, v)
			}
		}
	}
	return res
}
