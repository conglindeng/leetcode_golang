package middle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	mockHead := &ListNode{Next: head, Val: -1}
	//find mid ListNode
	slow, fast := mockHead, mockHead
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	//reverse
	cur := slow
	slow = slow.Next
	cur.Next = nil
	for slow != nil {
		next := slow.Next

		slow.Next = cur.Next
		cur.Next = slow

		slow = next
	}
	//merge
	slow, fast = mockHead.Next, cur.Next
	cur.Next = nil
	for slow != nil && fast != nil {
		slowNext := slow.Next
		fastNext := fast.Next

		slow.Next = fast
		fast.Next = slowNext

		slow = slowNext
		fast = fastNext
	}
}
