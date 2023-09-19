package middle

func Rob__(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res := 0
	first, second := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		tmp := max(first+nums[i], second)
		first = second
		second = tmp
	}
	res = max(res, second)
	first = 0
	second = 0
	for i := 1; i < len(nums); i++ {
		tmp := max(first+nums[i], second)
		first = second
		second = tmp
	}
	return max(res, second)
}
