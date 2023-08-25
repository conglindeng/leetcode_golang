package middle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers_445(l1 *ListNode, l2 *ListNode) *ListNode {
	return reverseLinkedList(addTwoNumbers(reverseLinkedList(l1), reverseLinkedList(l2)))
}

func reverseLinkedList(head *ListNode) *ListNode {
	mockHead := &ListNode{Val: -1}
	tmp := head

	for tmp != nil {
		next := tmp.Next
		tmp.Next = mockHead.Next
		mockHead.Next = tmp
		tmp = next
	}
	return mockHead.Next
}

func addTwoNumbers_445_with_stack(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := []int{}, []int{}
	tmp := l1
	for tmp != nil {
		stack1 = append(stack1, tmp.Val)
		tmp = tmp.Next
	}
	tmp = l2
	for tmp != nil {
		stack2 = append(stack2, tmp.Val)
		tmp = tmp.Next
	}
	mockHead := &ListNode{Val: -1}
	carry := 0
	for len(stack1) > 0 && len(stack2) > 0 {
		n1 := stack1[len(stack1)-1]
		n2 := stack2[len(stack2)-1]

		v := n1 + n2 + carry
		newNode := &ListNode{Val: v % 10, Next: mockHead.Next}
		mockHead.Next = newNode
		carry = v / 10

		stack1 = stack1[:len(stack1)-1]
		stack2 = stack2[:len(stack2)-1]
	}
	for len(stack1) > 0 {
		n1 := stack1[len(stack1)-1]

		v := n1 + carry
		newNode := &ListNode{Val: v % 10, Next: mockHead.Next}
		mockHead.Next = newNode
		carry = v / 10

		stack1 = stack1[:len(stack1)-1]
	}
	for len(stack2) > 0 {
		n2 := stack2[len(stack2)-1]

		v := n2 + carry
		newNode := &ListNode{Val: v % 10, Next: mockHead.Next}
		mockHead.Next = newNode
		carry = v / 10

		stack2 = stack2[:len(stack2)-1]
	}
	if carry > 0 {
		newNode := &ListNode{Val: carry, Next: mockHead.Next}
		mockHead.Next = newNode
	}
	return mockHead.Next
}
