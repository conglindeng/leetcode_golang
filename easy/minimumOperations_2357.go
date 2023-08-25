package easy

func minimumOperations(nums []int) int {
	count := map[int]struct{}{}
	for _, n := range nums {
		if n != 0 {
			count[n] = struct{}{}
		}
	}
	return len(count)
}
