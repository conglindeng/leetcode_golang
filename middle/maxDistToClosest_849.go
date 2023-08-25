package middle

func MaxDistToClosest(seats []int) int {
	pre := -1
	res := 0
	for i, v := range seats {
		if v == 1 {
			// update res
			if pre == -1 {
				res = i
			} else {
				res = max((i-pre)/2, res)
			}
			//update pre
			pre = i
		}
	}
	if pre < len(seats)-1 {
		res = max(res, len(seats)-1-pre)
	}
	return res
}
