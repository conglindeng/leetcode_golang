package easy

import (
	"sort"
)

func arithmeticTriplets(nums []int, diff int) int {
	count := 0
	for i := 0; i < len(nums)-2; i++ {
		if exsit(nums, nums[i]+diff) && exsit(nums, nums[i]+2*diff) {
			count++
		}

	}
	return count
}

func exsit(nums []int, num int) bool {
	m := sort.SearchInts(nums, num)
	return m != len(nums) && nums[m] == num
}
