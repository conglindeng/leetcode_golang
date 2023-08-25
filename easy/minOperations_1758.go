package easy

func MinOperations_1758(s string) int {
	var m = []byte{'0', '1'}
	getDiff := func(s string, idx int) int {
		c := 0
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] != m[idx%2] {
				c++
			}
			idx++
		}
		return c
	}
	return min(getDiff(s, 0), getDiff(s, 1))
}


func MinOperations_1758_(s string) int {
	cnt := 0
	for i, c := range s {
		if i%2 != int(c-'0') {
			cnt++
		}
	}
	return min(cnt, len(s)-cnt)
}
