package middle

import "unicode"

func letterCasePermutation_bfs(s string) []string {
	queue := []string{""}
	ans := []string{}
	for len(queue) > 0 {
		cur := queue[0]
		pos := len(cur)
		if pos == len(s) {
			ans = append(ans, cur)
			queue = queue[1:]
		} else {
			if unicode.IsLetter(rune(s[pos])) {
				queue = append(queue, cur+string(s[pos]^32))
			}
			queue[0] += string(s[pos])
		}
	}
	return ans
}

func letterCasePermutation_backtracking(s string) []string {
	ans := []string{}
	

	return ans
}
