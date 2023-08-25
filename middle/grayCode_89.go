package middle

func grayCode(n int) []int {
	l := 1 << n
	res := make([]int, 0)
	for i := 0; i < l; i++ {
		res = append(res, i^(i>>1))
	}
	return res
}
