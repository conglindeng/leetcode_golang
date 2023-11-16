package easy

func LongestAlternatingSubarray(nums []int, threshold int) int {
	//start is odd,
	i := 0
	res := 0
	l := 0
	for i < len(nums) {
		if nums[i] > threshold {
			l = 0

		} else if nums[i]%2 == 0 {
			if i-1 >= 0 && nums[i-1]%2 == 0 {
				l = 1
			} else {
				l += 1
			}
		} else {
			if i-1 >= 0 && nums[i-1] <= threshold && nums[i-1]%2 == 0 {
				l += 1
			} else {
				l = 0
			}
		}
		i++
		if l > res {
			res = l
		}
	}
	return res
}
