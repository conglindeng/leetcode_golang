package middle

import "sort"

func CanPartitionKSubsets(nums []int, k int) bool {
	l := len(nums)
	if l < k {
		return false
	}
	sort.Ints(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	left, right := 0, l-1
	target := sum / k
	for left < right {
		if nums[right] == target {
			right--
		} else if nums[right]+nums[left] == target {
			right--
			left++
		} else {
			return false
		}
	}
	return true
}
