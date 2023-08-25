package middle

func isRobotBounded(instructions string) bool {
	var direct = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	base := 0
	x, y := 0, 0
	for i := 0; i < len(instructions); i++ {
		c := instructions[i]
		if c == 'G' {
			x, y = x+direct[base][0], y+direct[base][1]
		} else if c == 'L' {
			base = (base + 4 - 1) % 4
		} else if c == 'R' {
			base = (base + 1) % 4
		}
	}
	return base != 0 || (x == 0 && y == 0)
}
