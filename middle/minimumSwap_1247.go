package middle

func MinimumSwap(s1 string, s2 string) int {
	if len(s1) != len(s2) {
		return -1
	}
	xy, yx := 0, 0
	for i := 0; i < len(s1); i++ {
		a, b := s1[i], s2[i]
		if a == 'x' && b == 'y' {
			xy++
		} else if b == 'x' && a == 'y' {
			yx++
		}
	}
	if (xy+yx)%2 == 1 {
		return -1
	}
	return xy/2 + yx/2 + xy%2 + yx%2
}
