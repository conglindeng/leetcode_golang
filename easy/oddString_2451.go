package easy

func oddString(words []string) string {
	differ := make([][]int, len(words))
	l := len(words[0])
	for i := 0; i < len(words); i++ {
		differ[i] = make([]int, l-1)
	}
	for j, s := range words {
		for i := 0; i < l-1; i++ {
			differ[j][i] = int(s[i+1] - s[i])
		}
	}
	for i := 0; i < l-1; i++ {
		for j := 1; j < len(differ); j++ {
			if differ[j][i] != differ[j-1][i] {
				if j > 1 {
					return words[j]
				} else {
					if differ[2][i] == differ[1][i] {
						return words[0]
					} else {
						return words[1]
					}
				}
			}
		}
	}
	return ""
}
