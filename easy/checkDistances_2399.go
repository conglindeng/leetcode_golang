package easy

func checkDistances(s string, distance []int) bool {
	position := map[int]int{}
	for i := 0; i < len(s); i++ {
		n := int(s[i] - 'a')
		if p, exist := position[n]; exist {
			if (i - p - 1) != distance[n] {
				return false
			}
		}
		position[n] = i
	}
	return true
}
