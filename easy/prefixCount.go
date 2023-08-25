package easy

import "strings"

func prefixCount(words []string, pref string) int {
	count := 0
	for _, s := range words {
		if strings.HasPrefix(s, pref) {
			count++
		}
	}
	return count
}
