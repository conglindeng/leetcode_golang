package middle

import (
	"math"
	"sort"
)

// [4,6,5,9,3,7]
// [0,0,2]
// [2,3,5]
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	res := make([]bool, len(l))
	for i := 0; i < len(l); i++ {
		res[i] = validate_(nums, l[i], r[i])
	}
	return res
}

func validate(nums []int, i, j int) bool {
	tmp := make([]int, 0)
	tmp = append(tmp, nums[i:j+1]...)
	sort.Ints(tmp)
	d := tmp[1] - tmp[0]
	for i := 2; i < len(tmp); i++ {
		if tmp[i]-tmp[i-1] != d {
			return false
		}
	}
	return true
}

func validate_(nums []int, i, j int) bool {
	mi, ma, l := math.MaxInt, math.MinInt, j-i+1
	for m := i; m <= j; m++ {
		if nums[m] > ma {
			ma = nums[m]
		}
		if nums[m] < mi {
			mi = nums[m]
		}
	}
	if ma == mi {
		return true
	}
	if (ma-mi)%(l-1) != 0 {
		return false
	}
	step := (ma - mi) / (l - 1)
	exist := make([]bool, l)
	for m := i; m <= j; m++ {
		cur := nums[m] - mi
		if cur%step != 0 {
			return false
		}
		exist[cur/step] = true
	}
	for _, v := range exist {
		if !v {
			return false
		}
	}
	return true
}
