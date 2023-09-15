package algorithm

import "strconv"

func calculate(s string) int {
	return -1
}

func Change2RPN(s string) []string {
	preIsOp, sign := true, 1
	stack := []string{}
	opStack := []string{}
	i, n := 0, len(s)
	for i < n {
		switch s[i] {
		case ' ':
			i++
		case '*':
			preIsOp = true
			// if len(opStack) > 0 && (opStack[len(opStack)-1] == "+" || opStack[len(opStack)-1] == "-") {
			// 	top := opStack[len(opStack)-1]
			// 	opStack[len(opStack)-1]="*"

			// 	opStack = append(opStack, top)
			// } else {
				opStack = append(opStack, "*")
			// }
			i++
		case '/':
			preIsOp = true
			// if len(opStack) > 0 && (opStack[len(opStack)-1] == "+" || opStack[len(opStack)-1] == "-") {
			// 	top := opStack[len(opStack)-1]
			// 	opStack[len(opStack)-1]="/"
			// 	opStack = append(opStack, top)
			// } else {
				opStack = append(opStack, "/")
			// }
			i++
		case '+':
			preIsOp = true
			if len(opStack) > 0 && (opStack[len(opStack)-1] == "*" || opStack[len(opStack)-1] == "/") {
				stack = append(stack, opStack[len(opStack)-1])
				opStack[len(opStack)-1] = "+"
			} else {
				opStack = append(opStack, "+")
			}
			i++
		case '-':
			if preIsOp {
				preIsOp = false
				sign = -1
			} else {
				preIsOp = true

				if len(opStack) > 0 && (opStack[len(opStack)-1] == "*" || opStack[len(opStack)-1] == "/") {
					stack = append(stack, opStack[len(opStack)-1])
					opStack[len(opStack)-1] = "-"
				} else {
					opStack = append(opStack, "-")
				}
			}
			i++
		case '(':
			preIsOp = true
			opStack = append(opStack, "(")
			i++
		case ')':
			opStack = append(opStack, ")")
			preIsOp = false
			i++
		default:
			cur := 0
			for ; i < n && s[i] >= '0' && s[i] <= '9'; i++ {
				cur = cur*10 + int(s[i]-'0')
			}
			stack = append(stack, strconv.Itoa(sign*cur))
			sign = 1
			preIsOp = false
		}
	}
	for len(opStack) > 0 {
		cur := opStack[len(opStack)-1]
		opStack = opStack[:len(opStack)-1]
		if cur == "(" || cur == ")" {
			continue
		} else {
			stack = append(stack, cur)
		}
	}
	return stack
}
