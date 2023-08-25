package middle

func CountVowelStrings(n int) int {
	count := []int{1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		for m := 0; m < 5; m++ {
			k := count[m]
			for n := m + 1; n < 5; n++ {
				k += count[n]
			}
			count[m] = k
		}
	}
	res := 0
	for _, n := range count {
		res += n
	}
	return res
}
