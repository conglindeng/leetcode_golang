package middle

func numPairsDivisibleBy60(time []int) int {
	count := map[int]int{}
	res := 0
	for _, v := range time {
		cur := v % 60
		res += count[(60-cur)%60]
		count[cur]++
	}
	return res
}
