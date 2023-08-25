package middle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveZeroSumSublists(head *ListNode) *ListNode {
	prefix := map[int]*ListNode{}
	mockHead := &ListNode{Next: head, Val: 0}
	sum := 0
	tmp := mockHead
	for ; tmp != nil; tmp = tmp.Next {
		sum += tmp.Val
		prefix[sum] = tmp
	}
	tmp = mockHead
	sum = 0
	for ; tmp != nil; tmp = tmp.Next {
		sum += tmp.Val
		tmp.Next = prefix[sum].Next
	}
	return mockHead.Next
}
