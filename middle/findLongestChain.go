package middle

import (
	"sort"
)

var (
	max_Length_Chain int
)

func FindLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		first := pairs[i]
		second := pairs[j]
		if first[0] != second[0] {
			return first[0] < second[0]
		} else {
			return first[1] < second[1]
		}
	})
	max_Length_Chain = 1
	for i := 0; i < len(pairs)-1; i++ {
		rescusionToFindMax(i, 1, pairs[i][1], pairs)
	}
	return max_Length_Chain
}

func rescusionToFindMax(index, length, pre int, pairs [][]int) {
	max_Length_Chain = max(max_Length_Chain, length)
	for i := index + 1; i < len(pairs); i++ {
		if pairs[i][0] > pre {
			rescusionToFindMax(i, length+1, pairs[i][1], pairs)
		}
	}
}

func findLongestChain_dp(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		first := pairs[i]
		second := pairs[j]
		if first[0] != second[0] {
			return first[0] < second[0]
		} else {
			return first[1] < second[1]
		}
	})
	res := 1
	dp := make([]int, len(pairs))
	for i, n := range pairs {
		dp[i] = 1
		for j, m := range pairs[:i] {
			if m[1] < n[0] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func findLongestChain_subsequence(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][0] > pairs[j][0] })
	res := make([]int, 0)
	for _, item := range pairs {
		i := sort.SearchInts(res, item[0])
		if i<len(res){
			res[i]=min(res[i],item[1])
		}else {
			res = append(res, item[1])
		}
	}
	return len(res)
}
