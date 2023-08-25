package easy

import "strings"

func strongPasswordCheckerII(password string) bool {
	const specialChar = "!@#$%^&*()-+"
	if len(password) < 8 {
		return false
	}
	lower, upper, digit, special := false, false, false, false
	for i := range password {
		if i > 0 && password[i] == password[i-1] {
			return false
		}
		c := byte(password[i])
		if c >= 'A' && c <= 'Z' {
			upper = true
		} else if c >= 'a' && c <= 'z' {
			lower = true
		} else if c >= '0' && c <= '9' {
			digit = true
		} else if strings.ContainsRune(specialChar, rune(c)) {
			special = true
		}
	}
	return lower && upper && digit && special
}
