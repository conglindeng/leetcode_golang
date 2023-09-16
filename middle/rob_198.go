package middle

func rob(nums []int) int {
	//dp[i]=max(dp[i-1],dp[i-2]+nums[i])
	first, second := 0, 0
	for i := 0; i < len(nums); i++ {
		tmp := max(second, first+nums[i])
		first = second
		second = tmp
	}
	return second
}
