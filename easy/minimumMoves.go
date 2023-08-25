package easy

func minimumMoves(s string) int {
	i, n := 0, len(s)
	res := 0
	for i < n {
		if s[i] == 'X' {
			res++
			i += 3
		} else {
			i++
		}
	}
	return res
}
