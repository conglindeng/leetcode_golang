package middle

func MovesToMakeZigzag(nums []int) int {
	cal := func(j int) int {
		res := 0
		for i := j; i < len(nums); i += 2 {
			a := 0
			if i > 0 {
				a = max(a, nums[i]-nums[i-1]+1)
			}
			if i+1 < len(nums) {
				a = max(a, nums[i]-nums[i+1]+1)
			}
			res += a
		}
		return res
	}
	return min(cal(0), cal(1))
}
