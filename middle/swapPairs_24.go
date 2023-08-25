package middle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	mockHead := &ListNode{Next: head}
	tmp := mockHead
	for tmp.Next != nil {
		next2 := tmp.Next.Next
		next := tmp.Next
		if next2 == nil {
			break
		}
		next.Next = next2.Next
		tmp.Next = next2
		next2.Next = next

		tmp = next
	}

	return mockHead.Next
}
