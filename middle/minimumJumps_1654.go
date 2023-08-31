package middle

type jumpInfo struct {
	pos, step int
	preIsLeft bool
}

// out of memory, from zero to x
func minimumJumps(forbidden []int, a int, b int, x int) int {
	block := map[int]bool{}
	maxVal := 0
	for _, v := range forbidden {
		block[v] = true
		maxVal = max(maxVal, v)
	}
	lower, upper := 0, max(maxVal+a, x)+b
	visit := map[int]bool{}
	visit[0] = true
	validate := func(pos, dire int) bool {
		if block[pos] {
			return false
		}
		if visit[pos*dire] {
			return false
		}
		if pos < lower || pos > upper {
			return false
		}
		return true
	}
	work := []jumpInfo{}
	work = append(work, jumpInfo{pos: 0, preIsLeft: true, step: 0})
	for len(work) > 0 {
		cur := work[0]
		work = work[1:]
		if cur.pos == x {
			return cur.step
		}
		if !cur.preIsLeft {
			if validate(cur.pos-b, -1) {
				work = append(work, jumpInfo{pos: cur.pos - b, preIsLeft: true, step: cur.step + 1})
				visit[cur.pos-b] = true
			}
		}
		if validate(cur.pos+a, 1) {
			work = append(work, jumpInfo{pos: cur.pos + a, preIsLeft: false, step: cur.step + 1})
			visit[cur.pos+a] = true
		}
	}
	return -1
}
