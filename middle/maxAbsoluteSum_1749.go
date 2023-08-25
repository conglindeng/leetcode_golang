package middle

func MaxAbsoluteSum(nums []int) int {
	positive, negative := 0, 0
	res := 0
	for _, v := range nums {
		positive += v
		negative += v
		res = max(res, max(intAbs(positive, 0), intAbs(negative, 0)))
		if negative > 0 {
			negative = 0
		}
		if positive < 0 {
			positive = 0
		}
	}
	return res
}
