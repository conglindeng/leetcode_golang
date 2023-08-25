package easy

import "sort"

func AnswerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	count := make([]int, len(nums))
	count[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		count[i] = count[i-1] + nums[i]
	}
	res := []int{}
	// for _, n := range queries {
	// 	res = append(res, sort.SearchInts(count, n+1))
	// }
	for _, n := range queries {
		idx := sort.Search(len(nums), func(i int) bool {
			return count[i] >= n+1
		})
		res = append(res, idx)
	}
	return res
}
