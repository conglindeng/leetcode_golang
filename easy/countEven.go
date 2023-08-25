package easy

func countEven(num int) int {
	count := 0
	for i := 1; i <= num; i++ {
		if isEven(i) {
			count++
		}
	}
	return count
}

func isEven(num int) bool {
	sum := 0
	for num > 0 {
		sum += num % 10
		num = num / 10
	}
	return sum%2 == 0
}