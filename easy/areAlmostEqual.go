package easy

// 4ms 1.9mb
func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 1 || len(s2) == 1 {
		return s1 == s2
	}
	i, j := -1, -1
	for m, c := range s1 {
		//c is rune
		if byte(c) != s2[m] {
			if i == -1 {
				i = m
			} else if j == -1 {
				j = m
			} else {
				return false
			}
		}
	}
	return j != -1 && i != -1 && s2[j] == s1[i] && s2[i] == s1[j]
}

// 0ms 1.9mb
func areAlmostEqual_(s1 string, s2 string) bool {
	i, j := -1, -1
	for m, c := range s1 {
		//c is rune
		if byte(c) != s2[m] {
			if i == -1 {
				i = m
			} else if j == -1 {
				j = m
			} else {
				return false
			}
		}
	}
	return (i < 0) || (j > 0 && s2[j] == s1[i] && s2[i] == s1[j])
}
