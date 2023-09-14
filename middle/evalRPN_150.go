package middle

import "strconv"

func EvalRPN(tokens []string) int {
	var doEval func(tokens []string, pos int) (int, int)
	doEval = func(tokens []string, pos int) (int, int) {
		if len(tokens) == 1 {
			x, _ := strconv.Atoi(tokens[0])
			return x, -1
		}
		cur := tokens[pos]
		switch cur {
		case "*":
			second, pos1 := doEval(tokens, pos-1)
			first, pos2 := doEval(tokens, pos1-1)
			return first * second, pos2
		case "/":
			second, pos1 := doEval(tokens, pos-1)
			first, pos2 := doEval(tokens, pos1-1)
			return first / second, pos2
		case "+":
			second, pos1 := doEval(tokens, pos-1)
			first, pos2 := doEval(tokens, pos1-1)
			return first + second, pos2
		case "-":
			second, pos1 := doEval(tokens, pos-1)
			first, pos2 := doEval(tokens, pos1-1)
			return first - second, pos2
		default:
			x1, _ := strconv.Atoi(tokens[pos])
			return x1, pos
		}
	}
	x, _ := doEval(tokens, len(tokens)-1)
	return x
}

func EvalRPN_forward(tokens []string) int {
	stack := []int{}
	for _, v := range tokens {
		if v[len(v)-1] <= '9' && v[len(v)-1] >= '0' {
			x, _ := strconv.Atoi(v)
			stack = append(stack, x)
		} else {
			second := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch v {
			case "*":
				stack[len(stack)-1] = stack[len(stack)-1] * second
			case "/":
				stack[len(stack)-1] = stack[len(stack)-1] / second
			case "+":
				stack[len(stack)-1] = stack[len(stack)-1] + second
			case "-":
				stack[len(stack)-1] = stack[len(stack)-1] - second
			}
		}
	}
	return stack[0]
}
