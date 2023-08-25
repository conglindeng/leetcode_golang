package middle

func maxChunksToSorted(arr []int) int {
	res, max := 0, 0
	for i, n := range arr {
		if n > max {
			max = n
		}
		if i == max {
			res++
		}
	}
	return res
}
