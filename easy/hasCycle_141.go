package easy

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for slow != nil && slow.Next != nil {
		slow = slow.Next
		if slow.Next == fast {
			return true
		}
		fast = fast.Next
		slow = slow.Next
	}
	return false
}
