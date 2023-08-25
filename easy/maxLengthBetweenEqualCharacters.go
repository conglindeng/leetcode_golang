package easy

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxLengthBetweenEqualCharacters(s string) int {
	index := map[rune]int{}
	res := -1
	for i, c := range s {
		m, ok := index[c]
		if !ok {
			index[c] = i
		} else {
			res = max(res, i-m-1)
		}
	}
	return res
}
