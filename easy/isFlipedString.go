package easy

import "strings"

/**
Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 (e.g.,"waterbottle" is a rotation of"erbottlewat"). Can you use only one call to the method that checks if one word is a substring of another?

Example 1:

Input: s1 = "waterbottle", s2 = "erbottlewat"
Output: True
**/

func IsFlipedString(s1 string, s2 string) bool {
	if len(s1) == 0 && len(s2) == 0 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}

	i, j := 0, 0
	for j < len(s1) {
		if s1[i] == s2[j] {
			temp := i
			m := 0
			for ; m < len(s2); m++ {
				if s1[temp] != s2[(j+m)%len(s2)] {
					break
				}
				temp++
			}
			if m == len(s2) {
				return true
			}
		}
		j++
	}
	return false
}

func isFlipedString(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s1+s1, s2)
}
