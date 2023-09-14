package middle

func Calculate(s string) int {
	nums := []int{0}
	ops := []byte{'+'}
	i, n := 0, len(s)
	for i < n {
		switch s[i] {
		case ' ':
			i++
		case '+':
			ops = append(ops, '+')
			i++
		case '-':
			ops = append(ops, '-')
			i++
		case '*':
			ops = append(ops, '*')
			i++
		case '/':
			ops = append(ops, '/')
			i++
		default:
			cur := 0
			for ; i < n && '0' <= s[i] && '9' >= s[i]; i++ {
				cur = cur*10 + int(s[i]-'0')
			}
			if len(ops) > 0 && (ops[len(ops)-1] == '/' || ops[len(ops)-1] == '*') {
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1]
				if op == '*' {
					nums[len(nums)-1] = nums[len(nums)-1] * cur
				} else {
					nums[len(nums)-1] = nums[len(nums)-1] / cur
				}
			} else {
				nums = append(nums, cur)
			}
		}
	}
	for i := 0; i < len(ops); i++ {
		op := ops[i]
		cur := nums[i]
		if op == '+' {
			nums[i+1] = cur + nums[i+1]
		} else {
			nums[i+1] = cur - nums[i+1]
		}
	}
	return nums[len(nums)-1]
}
