package difficult

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	mockHead := &ListNode{Val: -1}
	tmp := mockHead
	left, right := 0, len(lists)-1
	for right >= 0 && lists[right] == nil {
		right--
	}
	for left <= right {
		if lists[left] == nil {
			lists[left], lists[right] = lists[right], lists[left]
			left++
			right--
		} else {
			left++
		}
	}
	if right < 0 {
		return nil
	}
	lists = lists[:right+1]
	for len(lists) > 0 {
		idx := 0
		for i := 1; i < len(lists); i++ {
			if lists[i].Val < lists[idx].Val {
				idx = i
			}
		}

		next := lists[idx].Next
		lists[idx].Next = nil
		tmp.Next = lists[idx]
		tmp = tmp.Next

		if next == nil {
			lists[idx] = lists[len(lists)-1]
			lists = lists[:len(lists)-1]
		} else {
			lists[idx] = next
		}
	}
	return mockHead.Next
}

func mergeKLists_divide_conquer(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return lists[0]
	}
	if l == 2 {
		return mergeTwoList(lists[0], lists[1])
	}
	mid := l / 2
	left := mergeKLists_divide_conquer(lists[:mid])
	right := mergeKLists_divide_conquer(lists[mid:])
	return mergeTwoList(left, right)
}

func mergeTwoList(first, second *ListNode) *ListNode {
	mockHead := &ListNode{Val: -1}
	tmp := mockHead

	for first != nil && second != nil {
		if first.Val < second.Val {
			tmp.Next = first
			first = first.Next
		} else {
			tmp.Next = second
			second = second.Next
		}
		tmp = tmp.Next
	}
	if first != nil {
		tmp.Next = first
	}
	if second != nil {
		tmp.Next = second
	}
	return mockHead.Next
}
