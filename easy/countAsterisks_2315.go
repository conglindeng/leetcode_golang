package easy

func CountAsterisks(s string) int {
	res := 0
	flag := false
	for i := 0; i < len(s); i++ {
		c := byte(s[i])
		if c == '|' {
			flag = !flag
		} else if c == '*' {
			if !flag {
				res++
			}
		}
	}
	return res
}
