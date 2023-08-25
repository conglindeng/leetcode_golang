package middle

func minimumLength(s string) int {
	left, right := 0, len(s)-1
	for left < right && s[left] == s[right] {
		for left < right && s[left] == s[left+1] {
			left++
		}
		for left < right && s[right] == s[right-1] {
			right--
		}
		if left >= right {
			return 0
		}
		left++
		right--
	}
	return right - left + 1
}

func MinimumLength(s string) int {
	left, right := 0, len(s)-1
	for left < right && s[left] == s[right] {
		c := s[left]
		for left <= right && s[left] == c {
			left++
		}
		for left <= right && s[right] == c {
			right--
		}
	}
	return right - left + 1
}
