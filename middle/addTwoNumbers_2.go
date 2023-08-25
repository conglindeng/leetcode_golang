package middle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	tmp1, tmp2 := l1, l2
	mockHead := &ListNode{Val: -1}
	head := mockHead
	for tmp1 != nil && tmp2 != nil {
		v := carry + tmp1.Val + tmp2.Val
		head.Next = &ListNode{Val: v % 10}
		carry = v / 10
		tmp1 = tmp1.Next
		tmp2 = tmp2.Next
		head = head.Next
	}
	for tmp1 != nil {
		v := carry + tmp1.Val
		head.Next = &ListNode{Val: v % 10}
		carry = v / 10
		tmp1 = tmp1.Next
		head = head.Next
	}
	for tmp2 != nil {
		v := carry + tmp2.Val
		head.Next = &ListNode{Val: v % 10}
		carry = v / 10
		tmp2 = tmp2.Next
		head = head.Next
	}
	if carry != 0 {
		head.Next = &ListNode{Val: carry}
	}
	return mockHead.Next
}
