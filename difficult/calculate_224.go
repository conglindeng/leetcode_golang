package difficult

func Calculate(s string) int {
	ans := 0
	ops := []int{1}
	sign := 1
	n := len(s)
	for i := 0; i < n; {
		c := s[i]
		switch c {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && '9' >= s[i]; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
		}

	}
	return ans
}
