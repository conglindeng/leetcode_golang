package easy

func maximizeSum(nums []int, k int) int {
	m := nums[0]
	for _, v := range nums {
		if m < v {
			m = v
		}
	}
	return k * (m + m + k - 1) / 2
}
