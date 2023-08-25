package middle

import (
	"fmt"
	"sort"
)

func FindingUsersActiveMinutes(logs [][]int, k int) []int {
	count := map[int]int{}
	sort.Slice(logs, func(i, j int) bool {
		if logs[i][0] == logs[j][0] {
			return logs[i][1] < logs[j][1]
		} else {
			return logs[i][0] < logs[j][0]
		}
	})
	for i, m := range logs {
		if i > 0 {
			if m[0] == logs[i-1][0] && m[1] == logs[i-1][1] {
				continue
			} else {
				count[m[0]]++
			}
		} else {
			count[m[0]]++
		}
	}
	fmt.Println(count)
	res := make([]int, k)
	for _, v := range count {
		if v <= k {
			res[v-1]++
		}
	}
	return res
}
