package easy

func CalculateTax(brackets [][]int, income int) float64 {
	var res float64 = 0.0
	for i, num := range brackets {
		var pre int
		if i == 0 {
			pre = 0
		} else {
			pre = brackets[i-1][0]
		}
		var l int
		if income >= num[0] {
			l = num[0] - pre
		} else {
			l = income - pre
		}
		res = res + float64(l*num[1])*0.01
		if income < num[0] {
			break
		}
	}
	return res
}
