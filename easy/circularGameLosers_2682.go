package easy

func ÇircularGameLosers(n int, k int) []int {
	exist := map[int]bool{}
	exist[0] = true
	round := 1
	cur := 0
	for {
		cur = (cur + round*k) % n
		if exist[cur] {
			break
		}
		exist[cur] = true
		round++
	}
	res := []int{}
	for i := 0; i < n; i++ {
		if !exist[i] {
			res = append(res, i+1)
		}
	}
	return res
}
