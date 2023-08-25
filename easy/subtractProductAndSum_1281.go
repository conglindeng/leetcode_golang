package easy

func subtractProductAndSum(n int) int {
	sum, multi := 0, 1
	for n != 0 {
		k := n % 10
		n = n / 10
		sum += k
		multi *= k
	}
	return multi - sum
}
