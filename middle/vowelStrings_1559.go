package middle

func vowelStrings(words []string, queries [][]int) []int {
	var vowel = []byte{'a', 'e', 'i', 'o', 'u'}
	pre := make([]int, len(words)+1)
	for i := 0; i < len(words); i++ {
		if contains(vowel, words[i][0]) && contains(vowel, words[i][len(words[i])-1]) {
			pre[i+1] = pre[i] + 1
		} else {
			pre[i+1] = pre[i]
		}
	}
	res := make([]int, len(queries))
	for i, query := range queries {
		res[i] = pre[query[1]+1] - pre[query[0]]
	}
	return res
}

func contains(vowels []byte, b byte) bool {
	for _, by := range vowels {
		if by == b {
			return true
		}
	}
	return false
}
