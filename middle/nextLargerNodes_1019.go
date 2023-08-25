package middle

func nextLargerNodes(head *ListNode) []int {
	type ListNode struct {
		Val  int
		Next *ListNode
	}
	stack := make([][]int, 0)
	res := make([]int, 0)
	idx := -1
	cur := head
	for cur != nil {
		idx++
		res = append(res, 0)
		for len(stack) > 0 && stack[len(stack)-1][0] < cur.Val {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top[1]] = cur.Val
		}
		stack = append(stack, []int{cur.Val, idx})
		cur = cur.Next
	}
	return res
}
