package easy

func canFormArray(arr []int, pieces [][]int) bool {
	//every piece has the same order with arr
	// var getIdx func(int)int
	getIdx := func(target int) int {
		for i, n := range arr {
			if n == target {
				return i
			}
		}
		return -1
	}
	for _, nums := range pieces {
		start := nums[0]
		i := getIdx(start)
		if i == -1 {
			return false
		}
		for j := 0; j < len(nums); j++ {
			if i+j >= len(arr) {
				return false
			}
			if nums[j] != arr[i+j] {
				return false
			}
		}
	}
	return true
}
