package middle

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	res := int64(0)
	if cost1 < cost2 {
		cost2, cost1 = cost1, cost2
	}
	high := total / cost1
	for i := 0; i <= high; i++ {
		left := total - i*cost1
		res += int64(left/cost2) + 1
	}
	return res
}
