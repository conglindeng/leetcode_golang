package easy

import (
	"bytes"
)

func mergeAlternately(word1 string, word2 string) string {
	i, j := 0, 0
	var b bytes.Buffer
	for i < len(word1) && j < len(word2) {
		b.WriteByte(word1[i])
		b.WriteByte(word2[j])
		i++
		j++
	}
	for i < len(word1) {
		b.WriteByte(word1[i])
		i++
	}
	for j < len(word2) {
		b.WriteByte(word2[j])
		j++
	}
	return b.String()
}
