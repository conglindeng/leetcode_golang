package easy

import "strings"

func reorderSpaces(text string) string {
	words := strings.Fields(text)
	count := strings.Count(text, " ")
	l := len(words) - 1
	if l == 1 {
		return words[0] + strings.Repeat(" ", count)
	}
	return strings.Join(words, strings.Repeat(" ", count/l)) + strings.Repeat(" ", count%l)
}
