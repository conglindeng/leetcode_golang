package middle

func numMatchingSubseq(s string, words []string) int {
	count := 0
	flag := map[string]bool{}
	for _, v := range words {
		e, b := flag[v]
		if !b {
			exist := match(s, v)
			if exist {
				count++
			}
			flag[v] = exist
		}
		if e {
			count++
		}

	}
	return count
}

func match(s, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			i++
		}
	}
	return j == len(t)
}
