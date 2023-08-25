package middle

func ExpressiveWords(s string, words []string) int {
	count := 0
	for _, v := range words {
		if expand(s, v) {
			count++
		}
	}
	return count
}

func expand(s, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] != t[j] {
			return false
		}
		ch := s[i]
		count1 := 0
		for i < len(s) && s[i] == ch {
			count1++
			i++
		}
		count2 := 0
		for j < len(t) && t[j] == ch {
			count2++
			j++
		}
		if count1 < count2 || (count1 > count2 && count1 < 3) {
			return false
		}
	}
	return i == len(s) && j == len(t)
}
