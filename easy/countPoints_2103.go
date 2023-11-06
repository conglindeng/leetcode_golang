package easy

func countPoints(rings string) int {
	tmp := map[byte]int{}
	for i := 0; i < len(rings); i += 2 {
		color := rings[i]
		pos := rings[i+1]
		switch color {
		case 'R':
			tmp[pos] |= 1
		case 'G':
			tmp[pos] |= 1 << 1
		case 'B':
			tmp[pos] |= 1 << 2
		}
	}
	res := 0
	for _, v := range tmp {
		if v == 7 {
			res++
		}
	}
	return res
}
