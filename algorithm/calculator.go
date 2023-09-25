package algorithm

import "strconv"

func Calculate(s string) int {
	stack := []int{}
	tokens := Change2RPN(s)
	for _, v := range tokens {
		if v[len(v)-1] >= '0' && v[len(v)-1] <= '9' {
			x, _ := strconv.Atoi(v)
			stack = append(stack, x)
		} else {
			second := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch v {
			case "+":
				stack[len(stack)-1] += second
			case "-":
				stack[len(stack)-1] -= second
			case "*":
				stack[len(stack)-1] *= second
			case "/":
				stack[len(stack)-1] /= second
			}

		}

	}
	return stack[0]
}

// Reverse Polish Notation
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
			// opStack = append(opStack, ")")
			// pop op until top one is (
			for len(opStack) > 0 {
				cur := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]
				if cur == "(" {
					break
				}
				stack = append(stack, cur)
			}
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
