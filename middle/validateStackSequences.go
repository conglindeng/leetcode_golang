package middle

func ValidateStackSequences(pushed []int, popped []int) bool {
	//just
	stack := make([]int, 0)
	stack = append(stack, pushed[0])
	pushedIndex, popedIndex := 1, 0
	for {
		if pushedIndex == len(pushed) && popedIndex == len(popped) {
			break
		}
		for pushedIndex < len(pushed) && (len(stack) == 0 || stack[len(stack)-1] != popped[popedIndex]) {
			stack = append(stack, pushed[pushedIndex])
			pushedIndex++
		}
		if len(stack) > 0 && stack[len(stack)-1] == popped[popedIndex] {
			stack = stack[0 : len(stack)-1]
			popedIndex++
		} else {
			return false
		}
	}
	return true
}
