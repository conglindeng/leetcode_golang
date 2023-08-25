package easy

func minMaxGame(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	next := make([]int, len(nums)/2)
	for i := 0; i < len(next); i++ {
		if i%2 == 0 {
			next[i] = min(nums[2*i], nums[2*i+1])
		} else {
			next[i] = max(nums[2*i], nums[2*i+1])
		}
	}
	return minMaxGame(next)
}
