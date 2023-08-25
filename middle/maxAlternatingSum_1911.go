package middle

func MaxAlternatingSum(nums []int) int64 {
	odd, even := int64(0), int64(nums[0])
	for i := 1; i < len(nums); i++ {
		n := int64(nums[i])
		odd, even = int64Max(even-n, odd), int64Max(odd+n, even)
	}
	return int64(even)
}

func int64Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
