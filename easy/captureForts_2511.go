package easy

func CaptureForts(forts []int) int {
	pre := 0
	zeroCnt := 0
	res := 0
	for _, v := range forts {
		if v == 1 {
			if pre == -1 {
				res = max(zeroCnt, res)
			}
			zeroCnt = 0
			pre = 1
		} else if v == -1 {
			if pre == 1 {
				res = max(res, zeroCnt)
			}
			zeroCnt = 0
			pre = -1
		} else {
			zeroCnt++
		}
	}
	return res
}
