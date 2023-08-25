package middle

func ReinitializePermutation(n int) int {
	source := make([]int, 0, n)
	change := make([]int, 0, n)
	for i := 0; i < n; i++ {
		source = append(source, i)
		change = append(change, i)
	}
	step := 0
	for {
		step++
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				arr[i] = change[i/2]
			} else {
				arr[i] = change[n/2+(i-1)/2]
			}
		}
		change = arr
		if equal(source, change) {
			break
		}
	}
	return step
}

func equal(source, change []int) bool {
	for i := 0; i < len(source); i++ {
		if source[i] != change[i] {
			return false
		}
	}
	return true
}
