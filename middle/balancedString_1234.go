package middle

func BalancedString(s string) int {
	expected := len(s) / 4
	count := map[byte]int{}
	for _, c := range s {
		count[byte(c)]++
	}
	check := func(count map[byte]int) bool {
		if count['Q'] > expected ||
			count['W'] > expected ||
			count['E'] > expected ||
			count['R'] > expected {
			return false
		}
		return true
	}
	if check(count) {
		return 0
	}
	res := len(s)
	l, r := 0, 0
	for l < len(s) {
		for r < len(s) && !check(count) {
			count[s[r]]--
			r++
		}
		// case loop of r end with r==len(s)
		if !check(count) {
			break
		}
		res = min(res, r-l)
		count[s[l]]++
		l++
	}
	return res
}