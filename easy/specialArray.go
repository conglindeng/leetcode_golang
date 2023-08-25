package easy

import "sort"

func SpecialArray(nums []int) int {
	if len(nums) == 1 {
		if nums[0] >= 1 {
			return 1
		}
		return -1
	}
	sort.Ints(nums)
	// var left, right, mid int
	// left, right = 0, len(nums)
	// for left < right {
	// 	mid = (left + right) >> 1
	// 	count := len(nums) - getFirsGreaterOrEqual(nums, mid) + 1
	// 	if count == mid {
	// 		return count
	// 	} else if count < mid {
	// 		right = mid - 1
	// 	} else {
	// 		left = mid + 1
	// 	}
	// }

	for i := len(nums); i >= 0; i-- {
		count := len(nums) - getFirsGreaterOrEqual(nums, i)
		if count >= i {
			return count
		}
	}
	return -1
}

func getFirsGreaterOrEqual(nums []int, target int) int {
	var left, right, mid int
	left, right = 0, len(nums)
	for left < right {
		mid = (left + right) >> 1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func SpecialArray_(nums []int) int {
	if len(nums) == 1 {
		if nums[0] >= 1 {
			return 1
		}
		return -1
	}
	sort.Ints(nums)
	var left, right, mid int
	left, right = -1, len(nums)
	for left <= right {
		mid = (left + right) >> 1
		count := len(nums) - getFirsGreaterOrEqual(nums, mid)
		if count == mid {
			return count
		} else if count < mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
