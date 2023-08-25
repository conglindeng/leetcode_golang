package middle


// "(()(()))"
func scoreOfParentheses(s string) int {
	st := []int{0}
	for _, c := range s {
		if c == '(' {
			st = append(st, 0)
		} else {
			v := st[len(st)-1]
			st = st[:len(st)-1]
			st[len(st)-1] = max(2*v, 1)
		}
	}
	return st[0]
}

/**
Given a balanced parentheses string s, return the score of the string.

The score of a balanced parentheses string is based on the following rule:

"()" has score 1.
AB has score A + B, where A and B are balanced parentheses strings.
(A) has score 2 * A, where A is a balanced parentheses string.

**/

func scoreOfParentheses_(s string, index int) int {
	ans, bal := 0, 0
	for i, c := range s {
		if c == '(' {
			bal++
		} else {
			bal--
			if s[i-1] == '(' {
				ans += i << bal
			}
		}
	}
	return ans
}
