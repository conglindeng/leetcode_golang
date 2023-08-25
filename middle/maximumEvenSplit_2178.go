package middle

func MaximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 != 0 {
		return []int64{}
	}
	var evenSplit func(start, num int64) []int64
	evenSplit = func(start, num int64) []int64 {
		if start*2 >= num {
			return []int64{num}
		}
		res := []int64{start}
		tmp := evenSplit(start+2, num-start)
		res = append(res, tmp...)
		return res
	}
	return evenSplit(2, finalSum)
}

func MaximumEvenSplit_dp(finalSum int64) []int64 {
	if finalSum%2 != 0 {
		return []int64{}
	}
	start, num := int64(2), finalSum
	res := []int64{}
	for {
		if start*2 >= num {
			res = append(res, num)
			break
		} else {
			res = append(res, start)
			num -= start
			start += 2
		}
	}
	return res
}
