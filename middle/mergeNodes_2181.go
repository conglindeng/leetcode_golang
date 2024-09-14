package middle

func mergeNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mockHead := &ListNode{Val: -1}
	newCur := mockHead
	cur := head
	for cur != nil && cur.Val == 0 {
		cur = cur.Next
	}
	sum := 0
	for cur != nil {
		sum += cur.Val
		cur = cur.Next
		if cur != nil && cur.Val == 0 {
			newCur.Next = &ListNode{Val: sum}
			newCur = newCur.Next
			cur = cur.Next
			for cur != nil && cur.Val == 0 {
				cur = cur.Next
			}
			sum = 0
		}
	}
	if sum > 0 {
		newCur.Next = &ListNode{Val: sum}
		newCur = newCur.Next
	}
	return mockHead.Next
}

func mergeNodes_local(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mockHead := &ListNode{Val: -1}
	cur := head.Next
	mockCur := mockHead
	sum := 0
	for cur != nil {
		if cur.Val == 0 {
			mockCur.Next = &ListNode{Val: sum}
			mockCur = mockCur.Next
			sum = 0
		}
		sum += cur.Val
		cur = cur.Next
	}
	return mockHead.Next
}
