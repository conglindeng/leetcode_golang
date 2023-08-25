package middle

import "github.com/conglindeng/leetcode/mystruct"

func ReverseBetween(head *mystruct.ListNode, left int, right int) *mystruct.ListNode {
	if left >= right {
		return head
	}
	mockHead := &mystruct.ListNode{Next: head, Val: -1}
	pre := mockHead
	cur := head
	i := 0
	for ; i < left; i++ {
		pre = cur
		cur = cur.Next
	}
	last := cur
	for ; i < right; i++ {
		next := cur.Next

		cur.Next = pre.Next
		pre.Next = cur
		if i == right-1 {
			//the last one need process next carefully
			last.Next = next
		}
		cur = next
	}
	return mockHead.Next
}
