package middle

import m "github.com/conglindeng/leetcode/mystruct"

func MergeInBetween(list1 *m.ListNode, a int, b int, list2 *m.ListNode) *m.ListNode {
	tail := list2
	for tail.Next != nil {
		tail = tail.Next
	}
	mockHead := &m.ListNode{-1, list1}
	var aPre, bNext *m.ListNode
	cur := mockHead
	index := -1
	for index <= b {
		if index == a-1 {
			aPre = cur
		} else if index == b {
			bNext = cur.Next
		}
		index++
		cur = cur.Next
	}
	aPre.Next = list2
	tail.Next = bNext
	return mockHead.Next
}
