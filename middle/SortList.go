package middle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func SortList(head *ListNode) *ListNode {
	mockHead := &ListNode{}
	temp := head
	for temp != nil {
		next := temp.Next
		temp.Next = nil
		insert(mockHead, temp)
		temp = next
	}
	return mockHead.Next
}

func insert(head *ListNode, node *ListNode) {
	pre := head
	cur := head.Next
	for cur != nil && cur.Val < node.Val {
		pre = cur
		cur = cur.Next
	}
	node.Next = pre.Next
	pre.Next = node
}
