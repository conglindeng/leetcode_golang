package middle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, nums []int) int {
	numMap := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		numMap[n] = struct{}{}
	}
	temp := head
	count := 0
	for temp != nil {
		//find a sub-linkedList head
		if _, ok := numMap[temp.Val]; ok {
			for temp != nil {
				if _, ok := numMap[temp.Val]; !ok {
					break
				}
				temp = temp.Next
			}
			count++
		} else {
			temp = temp.Next
		}

	}
	return count
}

func numComponents_(head *ListNode, nums []int) int {
	numMap := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		numMap[n] = struct{}{}
	}
	count := 0
	for inset := false; head != nil; head = head.Next {
		if _, ok := numMap[head.Val]; !ok {
			inset = false
		} else if !inset {
			inset = true
			count++
		}
	}
	return count
}
