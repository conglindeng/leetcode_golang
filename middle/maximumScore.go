package middle

func maximumScore(a int, b int, c int) int {
	sum := a + b + c
	max := max(a, max(b, c))
	if sum < max*2 {
		return sum - max
	} else {
		return sum / 2
	}
}
