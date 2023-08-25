package easy

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	count := map[int]int{}
	for _, v := range items1 {
		count[v[0]] = v[1]
	}
	for _, v := range items2 {
		count[v[0]] += v[1]
	}
	res := make([][]int, 0)
	for k, v := range count {
		res = append(res, []int{k, v})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	return res
}
