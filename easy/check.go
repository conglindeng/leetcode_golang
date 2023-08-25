package easy

func Check(nums []int) bool {
	m := -1
	for i, v := range nums {
		if i > 0 && v < nums[i-1] {
			m = i
			break
		}
	}
	if m == -1 {
		return true
	}
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[(i+m+n)%n] < nums[(i+m+n-1)%n] {
			return false
		}
	}
	return true
}
