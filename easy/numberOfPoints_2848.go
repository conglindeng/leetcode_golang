package easy

import "sort"

func numberOfPoints(nums [][]int) int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i][0] == nums[j][0] {
			return nums[i][1] < nums[j][1]
		}
		return nums[i][0] < nums[j][0]
	})
	cnt := nums[0][1] - nums[0][0] + 1
	right := nums[0][1]
	for i := 1; i < len(nums); i++ {
		if nums[i][0] > right {
			cnt += nums[i][1] - nums[i][0] + 1
		} else if nums[i][1] > right {
			cnt += nums[i][1] - right
		}
		right = max(right, nums[i][1])
	}
	return cnt
}
