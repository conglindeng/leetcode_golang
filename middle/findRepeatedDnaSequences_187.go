package middle

func FindRepeatedDnaSequences(s string) []string {
	count := map[string]int{}
	res := []string{}
	for i := 0; i < len(s)-9; i++ {
		tmp := s[i : i+10]
		if count[tmp] == 1 {
			res = append(res, tmp)
		}
		count[tmp] += 1
	}

	return res
}

func FindRepeatedDnaSequences_sliding_window(s string) []string {
	if len(s) < 11 {
		return nil
	}
	count := map[int]int{}
	val := 0
	for i := 0; i < 9; i++ {
		val = val<<2 | mapByteToInt(s[i])
	}
	res := []string{}
	for i := 9; i < len(s); i++ {
		val = (val<<2 | mapByteToInt(s[i])) & ((1 << 20) - 1)
		count[val]++
		if count[val] == 2 {
			res = append(res, s[i-9:i+1])
		}
	}
	return res
}

func mapByteToInt(b byte) int {
	switch b {
	case 'A':
		return 0
	case 'C':
		return 1
	case 'G':
		return 2
	default:
		return 3
	}

}
