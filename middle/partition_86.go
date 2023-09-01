package middle

func partition(head *ListNode, x int) *ListNode {
	greaterHead := &ListNode{}
	lessHead := &ListNode{}
	greaterTemp := greaterHead
	lessTemp := lessHead

	tmp := head
	for tmp != nil {
		next := tmp.Next
		tmp.Next = nil
		if tmp.Val < x {
			greaterTemp.Next = tmp
			greaterTemp = greaterTemp.Next
		} else {
			lessTemp.Next = tmp
			lessTemp = lessTemp.Next
		}
		tmp = next
	}
	greaterTemp.Next = lessHead.Next
	return greaterHead.Next
}
