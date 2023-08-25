package middle

func GetSubarrayBeauty(nums []int, k int, x int) []int {
	n := len(nums)
	res := make([]int, 0)
	cnt := map[int]int{}
	for i := 0; i < k-1; i++ {
		if nums[i] < 0 {
			cnt[nums[i]]++
		}
	}
	for i := k - 1; i < n; i++ {
		if nums[i] < 0 {
			cnt[nums[i]]++
		}
		m := 0
		for j := -50; j < 0; j++ {
			m += cnt[j]
			if m >= x {
				res = append(res, j)
				break
			}
		}
		if m < x {
			res = append(res, 0)
		}
		if nums[i-k+1] < 0 {
			cnt[nums[i-k+1]]--
		}
	}
	return res
}
