package middle

import "sort"

func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int, 0)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			h, g := j+1, n-1
			for h < g {
				sum := nums[i] + nums[j] + nums[h] + nums[g]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[h], nums[g]})
					h++
					g--
					for h < g && nums[h] == nums[h-1] {
						h++
					}
					for h < g && nums[g] == nums[g+1] {
						g--
					}
				} else if sum < target {
					h++
				} else {
					g--
				}
			}
		}
	}
	return res
}
