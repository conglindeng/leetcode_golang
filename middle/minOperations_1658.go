package middle

//  [6016,5483,541,4325,8149,3515,7865,2209,9623,9763,4052,6540,2123,2074,765,7520,4941,5290,5868,
//   6150,6006,6077,2856,7826,9119]
// 31841

func MinOperations_(nums []int, x int) int {
	var do func(nums []int, left, right, x int) (bool, int)
	do = func(nums []int, left, right, x int) (bool, int) {
		if x == 0 {
			return true, 0
		}
		if left > right {
			return false, -1
		}
		if left == right && nums[left] == x {
			return true, 1
		}
		if nums[left] > x && nums[right] > x {
			return false, -1
		}
		if nums[left] <= x && nums[right] > x {
			b, i := do(nums, left+1, right, x-nums[left])
			if b {
				return b, i + 1
			}
		} else if nums[left] > x && nums[right] <= x {
			b, i := do(nums, left, right-1, x-nums[right])
			if b {
				return b, i + 1
			}
		} else if nums[left] <= x && nums[right] <= x {
			b, i := do(nums, left+1, right, x-nums[left])
			b2, i2 := do(nums, left, right-1, x-nums[right])
			if b && b2 {
				return true, min(i, i2) + 1
			} else if b && !b2 {
				return true, i + 1
			} else if !b && b2 {
				return true, i2 + 1
			} else {
				return false, -1
			}
		}
		return false, -1
	}
	bool, count := do(nums, 0, len(nums)-1, x)
	if !bool {
		return -1
	}
	return count
}

func MinOperations_prefix_suffix(nums []int, x int) int {
	pre := make([]int, len(nums)+1)
	suf := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		pre[i+1] = pre[i] + nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		suf[i] = suf[i+1] + nums[i]
	}
	left,right:=0,len(nums)
	res:=-1
	for left<right{
		//todo:dcl
	}
	return res
}
