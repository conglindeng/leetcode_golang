package easy

func MaxProduct(nums []int) int {
	max, secondMax := 0, 0
	for _, item := range nums {
		if item > secondMax {
			if item > max {
				max = item
			} else {
				secondMax = item
			}
		}
	}
	return (max - 1) * (secondMax - 1)
}
