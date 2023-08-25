package easy

func SecondHighest(s string) int {
	first, second := -1, -1
	for i := 0; i < len(s); i++ {
		if s[i] <= '9' && s[i] >= '0' {
			n := int(s[i] - '0')
			if n > first {
				second, first = first, n
			} else if second < n && n < first {
				second = n
			}
		}
	}
	return second
}