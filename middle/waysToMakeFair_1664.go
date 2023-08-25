package middle

//[2,1,6,4]

func waysToMakeFair(nums []int) int {
	rightDiff := make([]int, len(nums)+1)
	leftDiff := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			leftDiff[i+1] = nums[i] + leftDiff[i]
		} else {
			leftDiff[i+1] = -nums[i] + leftDiff[i]
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i%2 == 0 {
			rightDiff[i] = nums[i] + rightDiff[i+1]
		} else {
			rightDiff[i] = -nums[i] + rightDiff[i+1]
		}
	}
	res := 0
	for i := len(nums) - 1; i > 0; i-- {
		if leftDiff[i] == rightDiff[i+1] {
			res++
		}
	}
	return res
}
