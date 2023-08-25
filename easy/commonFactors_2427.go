package easy

func commonFactors(a int, b int) int {
	mini := min(b, a)
	thsMost := max(b, a)
	up := max(mini, (thsMost+1)/2)
	res := 1
	for i := 2; i <= up; i++ {
		if a%i == 0 && b%i == 0 {
			res++
		}
	}
	return res
}
