package middle

func minimumDeletions(s string) int {
	rightA := 0
	for _, c := range s {
		if c == 'a' {
			rightA++
		}
	}
	res := rightA
	leftB := 0
	for _, c := range s {
		if c == 'a' {
			rightA--
		} else {
			leftB++
		}
		res = min(res, leftB+rightA)
	}
	return res
}
