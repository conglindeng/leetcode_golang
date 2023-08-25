package easy

import "github.com/conglindeng/leetcode/mystruct"

func RemoveNthFromEnd(head *mystruct.ListNode, n int) *mystruct.ListNode {
	fast := head
	mockHead := &mystruct.ListNode{Next: head, Val: -1}
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	pre := mockHead
	slow := head
	for fast != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
	}
	pre.Next = pre.Next.Next
	return mockHead.Next
}
