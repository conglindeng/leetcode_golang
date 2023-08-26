package easy

import "strconv"

func summaryRanges(nums []int) []string {
	left := 0
	res := []string{}
	for left < len(nums) {
		right := left + 1
		for right < len(nums) && nums[right-1]+1 == nums[right] {
			right++
		}
		if left == right-1 {
			res = append(res, strconv.Itoa(nums[left]))
		} else {
			res = append(res, strconv.Itoa(nums[left])+"->"+strconv.Itoa(nums[right-1]))
		}
		left = right
	}
	return res
}
