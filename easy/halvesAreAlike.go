package easy

import "strings"

func halvesAreAlike(s string) bool {
	// var vowels = []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	left, right := 0, 0
	for i, c := range s {
		// if contains(vowels, s[i]) {
		if strings.ContainsRune("aeiouAEIOU", c) {
			if i < len(s)/2 {
				left++
			} else {
				right++
			}
		}
	}
	return left == right
}

func contains(runes []byte, cur byte) bool {
	for _, v := range runes {
		if v == cur {
			return true
		}
	}
	return false
}
