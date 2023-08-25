package easy

func averageValue(nums []int) int {
	sum, count := 0, 0
	for _, n := range nums {
		if n%6 == 0 {
			sum += n
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return sum / count
}
