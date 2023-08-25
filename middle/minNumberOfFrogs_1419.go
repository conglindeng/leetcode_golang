package middle

func minNumberOfFrogs(croakOfFrogs string) int {
	count := map[byte]int{}
	res := 0
	workCount := 0
	for _, c := range croakOfFrogs {
		if c == 'c' {
			workCount++
			res = max(res, workCount)
		} else if c == 'r' {
			count['c']--
			if count['c'] == -1 {
				return -1
			}
		} else if c == 'o' {
			count['r']--
			if count['r'] == -1 {
				return -1
			}
		} else if c == 'a' {
			count['o']--
			if count['o'] == -1 {
				return -1
			}
		} else if c == 'k' {
			count['a']--
			if count['a'] == -1 {
				return -1
			}
			workCount--
		} else {
			return -1
		}
		count[byte(c)]++
	}
	return res
}
