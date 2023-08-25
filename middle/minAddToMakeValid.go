package middle

import (
	"github.com/conglindeng/leetcode/mystruct"
)

func MinAddToMakeValid(s string) int {
	stack := mystruct.Stack_Init()
	res := 0
	for _, c := range s {
		if c == '(' {
			stack.Push(c)
		} else {
			if stack.IsEmpty() {
				res++
			} else if c, _ := stack.Peek(); c.(rune) == '(' {
				stack.Pop()
			} else {
				res++
			}
		}
	}
	res += stack.Size()
	return res
}
