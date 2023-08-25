package middle

import "math"

func NumTimesAllBlue(flips []int) int {
	res, max := 0, 0
	for i := 0; i < len(flips); i++ {
		if flips[i] > max {
			max = flips[i] - 1
		}
		if max == i {
			res++
		}
	}
	return res
}

func NumTimesAllBlue_(flips []int) int {
	n := len(flips)
	nums := make([]uint32, (n+31)/32)
	res := 0
	for i, f := range flips {
		m, n := (f-1)/32, (f-1)%32
		nums[m] = nums[m] ^ (1 << n)
		if validateAllBlue(nums, i) {
			res++
		}
	}
	return res
}

func validateAllBlue(nums []uint32, i int) bool {
	m, n := i/32, i%32
	for k := 0; k < m; k++ {
		if nums[k] != math.MaxUint32 {
			return false
		}
	}
	return nums[m] == (1<<(n+1) - 1)
}
