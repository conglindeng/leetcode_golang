package easy

func arraySign(nums []int) int {
	negetiveCount := 0
	for _, n := range nums {
		if n == 0 {
			return 0
		} else if n < 0 {
			negetiveCount++
		}
	}
	if negetiveCount%2 == 0 {
		return 1
	} else {
		return -1
	}
}

func arraySign_(nums []int) int {
	sign := 1
	for _, n := range nums {
		if n == 0 {
			return 0
		}
		if n < 0 {
			sign = -sign
		}
	}
	return sign
}
