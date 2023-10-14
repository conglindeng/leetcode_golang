package easy

import "strconv"

func findTheArrayConcVal(nums []int) int64 {
	res := int64(0)
	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		var s string
		if i != j {
			s = strconv.Itoa(nums[i]) + strconv.Itoa(nums[j])
		} else {
			s = strconv.Itoa(nums[i])
		}
		v, _ := strconv.ParseInt(s, 10, 64)
		res += v
	}
	return res
}
