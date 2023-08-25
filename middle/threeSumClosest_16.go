package middle

import "sort"

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 1; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				cur := nums[i] + nums[j] + nums[k]
				if intAbs(cur, target) < intAbs(res, target) {
					res = cur
				}
			}
		}
	}
	return res
}

func intAbs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func threeSumClosest_double_pointer(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2; i++ {
		j, k := i+1, n-1
		for j < k {
			x := nums[i] + nums[j] + nums[k]
			if abs(x-target) < abs(res-target) {
				res = x
			}
			if x > target {
				k--
			} else {
				j++
			}
		}
	}
	return res
}
