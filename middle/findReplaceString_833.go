package middle

import "bytes"

func FindReplaceString(s string, indices []int, sources []string, targets []string) string {
	var buffer bytes.Buffer
	index := map[int]int{}
	for j, v := range indices {
		index[v] = j
	}
	for i := 0; i < len(s); i++ {
		if _, exist := index[i]; exist {
			m := index[i]
			if matchSource(s, sources[m], i) {
				buffer.WriteString(targets[m])
				i = i + len(sources[m]) - 1
			} else {
				buffer.WriteByte(s[i])
			}

		} else {
			buffer.WriteByte(s[i])
		}
	}
	return buffer.String()
}

func matchSource(s, source string, start int) bool {
	for i := 0; i < len(source) && start+i < len(s); i++ {
		if source[i] != s[start+i] {
			return false
		}
	}
	return true
}
