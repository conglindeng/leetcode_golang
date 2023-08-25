package easy

func pivotInteger(n int) int {
	for i := 1; i <= n; i++ {
		preSum := i * (i + 1) / 2
		afterSum := (n - i + 1) * (i + n) / 2
		if preSum == afterSum {
			return i
		}
	}
	return -1
}
