package easy

import "sort"

func findMaxK(nums []int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	if nums[left] > 0 || nums[right] < 0 {
		return -1
	}
	for left < right {
		if nums[left] == nums[right] {
			return nums[right]
		} else if -nums[left] < nums[right] {
			right--
		} else {
			left++
		}
	}
	return -1
}
