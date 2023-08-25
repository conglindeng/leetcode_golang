package middle

func checkPalindromeFormation(a string, b string) bool {
	return check(a, b) || check(b, a)
}

func check(a, b string) bool {
	left, right := 0, len(a)-1
	for left < right && a[left] == b[right] {
		left++
		right--
	}
	if left >= right {
		return true
	}
	return isPalindrome(b[left:right+1]) || isPalindrome(a[left:right+1])
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j && s[i] == s[j] {
		i++
		j--
	}
	return i >= j
}
