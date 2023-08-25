package easy

import (
	"sort"
)

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	sum := 0
	for _,item := range arr[n/20 : 19*n/20] {
		sum += item
	}
	return float64(sum*10) / float64(n*9)
}
