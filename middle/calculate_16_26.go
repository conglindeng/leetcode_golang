package middle

import "strings"

func calculate(s string) int {
	//two stack, one is used to save number and another
	nums := make([]int, 0)
	ops := make([]byte, 0)
	n := 0
	s = strings.ReplaceAll(s, " ", "")
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '/' || c == '*' || c == '+' || c == '-' {
			if c == '-' || c == '+' {
				ops = append(ops, c)
				nums = append(nums, n)
				n = 0
			} else {
				second := 0
				j := i + 1
				for ; j < len(s) && s[j] >= '0' && s[j] <= '9'; j++ {
					second = second*10 + int(s[j]-'0')
				}
				if c == '*' {
					n = n * second
				} else {
					n = n / second
				}
				i = j - 1
			}
		} else if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		}
	}
	nums = append(nums, n)
	for len(ops) > 0 {
		op := ops[0]
		ops = ops[1:]
		second := nums[1]
		first := nums[0]
		nums = nums[1:]
		nums[0] = doCalculate(op, first, second)
	}
	return nums[0]
}

func doCalculate(op byte, first, second int) int {
	if op == '*' {
		return first * second
	} else if op == '/' {
		return first / second
	} else if op == '+' {
		return first + second
	} else {
		return first - second
	}
}

func Calculate_preop(s string) int {
	nums := make([]int, 0)
	n := 0
	pre := '+'
	s = strings.ReplaceAll(s, " ", "")
	for k, c := range s {
		digist := c <= '9' && c >= '0'
		if digist {
			n = n*10 + int(c-'0')
		}
		if !digist || k == len(s)-1 {
			if pre == '/' {
				nums[len(nums)-1] = nums[len(nums)-1] / n
			} else if pre == '*' {
				nums[len(nums)-1] = nums[len(nums)-1] * n
			} else if pre == '-' {
				nums = append(nums, -n)
			} else {
				nums = append(nums, n)
			}
			pre = c
			n = 0
		}
	}
	res := n
	for _, m := range nums {
		res += m
	}
	return res
}
