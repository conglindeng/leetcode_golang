package easy

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	s := strings.Split(sentence, " ")
	for i := 0; i < len(s); i++ {
		if strings.HasPrefix(s[i], searchWord) {
			return i
		}
	}
	return -1
}
