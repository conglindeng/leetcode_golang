package easy

func finalValueAfterOperations(operations []string) int {
	res := 0
	for _, v := range operations {
		if v == "++X" || v == "X++" {
			res++
		} else {
			res--
		}
	}
	return res
}
