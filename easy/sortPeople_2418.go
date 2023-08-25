package easy

import "sort"

func sortPeople(names []string, heights []int) []string {
	n := len(heights)
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		temp[i] = i
	}
	sort.Slice(temp, func(i, j int) bool {
		return heights[temp[i]] > heights[temp[j]]
	})
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = names[temp[i]]
	}
	return res
}
