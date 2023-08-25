package middle

func SortList_Optimize(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMid(head)

	return merge(head, mid)
}

func findMid(head *ListNode) *ListNode {
	mockHead := &ListNode{}
	mockHead.Next = head
	pre := mockHead
	fast := head
	slow := head
	for fast != nil {
		if fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		pre = slow
		slow = slow.Next
	}
	pre.Next = nil
	//断掉前后关系
	return slow
}

func merge(head1 *ListNode, head2 *ListNode) *ListNode {
	one := SortList_Optimize(head1)
	second := SortList_Optimize(head2)
	mockHead := &ListNode{}
	temp := mockHead
	for one != nil && second != nil {
		if one.Val > second.Val {
			temp.Next = second
			second = second.Next
		} else {
			temp.Next = one
			one = one.Next
		}
		temp = temp.Next
	}
	for one != nil {
		temp.Next = one
		one = one.Next
		temp=temp.Next
	}
	for second != nil {
		temp.Next = second
		second = second.Next
		temp=temp.Next
	}
	return mockHead.Next
}