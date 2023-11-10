package easy

func FindTheLongestBalancedSubstring(s string) int {
	res := 0
	for i := 1; i < len(s); i++ {
		if s[i] == '1' && s[i-1] == '0' {
			left, right := i-1, i
			for left >= 0 && right < len(s) && s[left] == '0' && s[right] == '1' {
				left--
				right++
			}
			res = max(res, right-left-1)
		}
	}
	return res
}

func findTheLongestBalancedSubstring(s string) int {
	count0, count1 := 0, 0
	res := 0
	for i, v := range s {
		if v == '0' {
			if i > 0 && s[i-1] == '1' {
				count0 = 1
				count1 = 0
			} else {
				count0++
			}
		} else {
			count1++
		}
		res = max(res, 2*min(count0, count1))
	}
	return res
}
