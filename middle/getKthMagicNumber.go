package middle

import (
	"container/heap"

	"github.com/conglindeng/leetcode/mystruct"
)

func GetKthMagicNumber(k int) int {
	factor := []int{3, 5, 7}
	visited := map[int]struct{}{1: {}}
	hp := mystruct.Heap_Init()
	hp.Push(1)
	//todo: why this
	for {
		cur := heap.Pop(hp).(int)
		if k == 1 {
			return cur
		}
		k--
		for _, v := range factor {
			if _, ok := visited[v*cur]; !ok {
				visited[v*cur] = struct{}{}
				heap.Push(hp, v*cur)
			}
		}
	}
}

func getKthMagicNumber_dp(k int) int {
	dp := make([]int, k+1)
	dp[1] = 1
	p3, p5, p7 := 1, 1, 1
	for i := 2; i <= k; i++ {
		dp[i] = min(min(dp[p3]*3, dp[p5]*5), dp[p7]*7)
		if dp[i] == dp[p3]*3 {
			p3++
		}
		if dp[i] == dp[p5]*5 {
			p5++
		}
		if dp[i] == dp[p7]*7 {
			p7++
		}
	}
	return dp[k]
}
