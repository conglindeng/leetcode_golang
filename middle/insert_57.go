package middle

import "sort"

func insert_57(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
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

func insert_57_again(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	left, right := newInterval[0], newInterval[1]
	merged := false
	for _, v := range intervals {
		if v[0] > right {
			if !merged {
				res = append(res, []int{left, right})
				merged = true
			}
			res = append(res, v)
		} else if v[1] < left {
			res = append(res, v)
		} else {
			left = min(left, v[0])
			right = max(right, v[1])
		}
	}
	if !merged {
		res = append(res, []int{left, right})
	}
	return res
}
