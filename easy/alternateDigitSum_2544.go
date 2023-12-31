package easy

func AlternateDigitSum(n int) int {
	nums := make([]int, 0)
	for n != 0 {
		nums = append(nums, n%10)
		n = n / 10
	}
	flag := 1
	res := 0
	for i := len(nums) - 1; i >= 0; i-- {
		res += nums[i] * flag
		flag = -flag
	}
	return res
}
