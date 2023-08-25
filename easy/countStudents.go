package easy

func CountStudents(students []int, sandwiches []int) int {
	s1 := 0
	for _, n := range students {
		s1 += n
	}
	s0 := len(students) - s1
	for _, m := range sandwiches {
		if m == 0 && s0 > 0 {
			s0--
		} else if m == 1 && s1 > 0 {
			s1--
		} else {
			break
		}
	}
	return s0 + s1
}
