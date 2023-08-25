package easy

func getLucky(s string, k int) int {
	sum := 0
	for i := range s {
		n := int(s[i] - 'a' + 1)
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
	}
	var cac func(sum, k int) int
	cac = func(sum, k int) int {
		if k == 1 || sum < 10 {
			return sum
		}
		temp := 0
		for sum > 0 {
			temp += sum % 10
			sum = sum / 10
		}
		return cac(temp, k-1)
	}
	return cac(sum, k)
}
