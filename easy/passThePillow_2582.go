package easy

func passThePillow(n int, time int) int {
	i := time % (n * 2)
	if i < n {
		return i
	}
	return 2*n - i
}

