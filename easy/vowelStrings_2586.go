package easy

func vowelStrings(words []string, left int, right int) int {
	var original = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	res := 0
	for left <= right {
		cur := words[left]
		if original[cur[0]] && original[cur[len(cur)-1]] {
			res++
		}
		left++
	}
	return res
}
