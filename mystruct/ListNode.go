// 删除倒数第N个节点
// 反转链表
package mystruct

type ListNode struct {
	Val  int
	Next *ListNode
}

func ChangeArray2ListNode(nums []int, index int) *ListNode {
	mockHead := &ListNode{Val: -1}
	cur := mockHead
	for _, n := range nums {
		cur.Next = &ListNode{Val: n}
		cur = cur.Next
	}
	if index > 0 {
		tmp := mockHead.Next
		for i := 0; i < index; i++ {
			tmp = tmp.Next
		}
		cur.Next = tmp
	}
	return mockHead.Next
}
