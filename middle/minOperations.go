package middle

func MinOperations(boxes string) []int {
	n := len(boxes)
	left := int(boxes[0] - '0')
	count := 0
	right := 0
	for i := 1; i < n; i++ {
		if boxes[i] == '1' {
			right++
			count += i
		}
	}
	res := make([]int, n)
	res[0] = count
	for i := 1; i < n; i++ {
		count += left - right
		if boxes[i] == '1' {
			right--
			left++
		}
		res[i] = count
	}
	return res
}
