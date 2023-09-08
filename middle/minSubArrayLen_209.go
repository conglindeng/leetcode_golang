package middle

func minSubArrayLen(target int, nums []int) int {
	i, j := 0, 0
	sum := 0
	res := len(nums) + 1
	for j < len(nums) {
		sum += nums[j]
		for sum >= target {
			res = min(res, j-i+1)
			sum -= nums[i]
			i++
		}
		j++
	}

	if res == len(nums)+1 {
		return 0
	}
	return res
}
