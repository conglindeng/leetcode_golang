package easy

func shuffle(nums []int, n int) []int {
	res := make([]int, len(nums))
	mid := len(nums) / 2
	for i, num := range nums {
		if i >= mid {
			res[(i-mid)*2+1] = num
		} else {
			res[i*2] = num
		}
	}
	return res
}
