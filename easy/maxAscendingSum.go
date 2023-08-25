package easy

func maxAscendingSum(nums []int) int {
	res := 0
	temp := 0
	for i, n := range nums {
		if i == 0 {
			temp += n
		} else {
			if nums[i-1] >= n {
				res = max(res, temp)
				temp = 0
			} else {
				temp += n
			}
		}
	}
	return max(res, temp)
}
