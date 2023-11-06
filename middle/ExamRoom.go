package middle

// type Node struct {
// 	val  int
// 	next *Node
// }

// type ExamRoom struct {
// 	n          int
// 	head, tail *Node
// }

// func ExamRoom_init(n int) ExamRoom {
// 	// func Constructor(n int) ExamRoom {
// 	return ExamRoom{n: n - 1}
// }

// func (r *ExamRoom) Seat() int {
// 	var location int
// 	if r.head == nil {
// 		node := &Node{val: 0}
// 		r.head = node
// 		r.tail = node
// 		location = 0
// 	} else if r.head == r.tail {
// 		cur := r.head.val
// 		if cur > r.n-cur {
// 			r.head = &Node{val: 0, next: r.tail}
// 			location = 0
// 		} else {
// 			r.tail = &Node{val: r.n}
// 			r.head.next = r.tail
// 			location = r.n
// 		}
// 	} else {
// 		maxZoo := r.head.val
// 		cur := r.head
// 		var pre *Node
// 		for cur != nil {
// 			var distance int
// 			if cur.next != nil {
// 				distance = (cur.next.val - cur.val) / 2
// 			} else {
// 				distance = r.n - cur.val
// 			}
// 			if distance > maxZoo {
// 				pre = cur
// 				maxZoo = distance
// 			}
// 			cur = cur.next
// 		}
// 		if pre != nil {
// 			node := &Node{val: pre.val + maxZoo, next: pre.next}
// 			pre.next = node
// 			location = node.val
// 		} else {
// 			newHead := &Node{val: 0, next: r.head}
// 			r.head = newHead
// 			location = newHead.val
// 		}
// 	}
// 	return location
// }

// func (r *ExamRoom) Leave(p int) {
// 	mockPre := &Node{val: -1, next: r.head}
// 	pre := mockPre
// 	cur := mockPre
// 	for cur.next != nil {
// 		if cur.next.val == p {
// 			pre = cur
// 			break
// 		}
// 		cur = cur.next
// 	}
// 	if pre.next == r.head {
// 		if r.head == r.tail {
// 			r.head = nil
// 			r.tail = nil
// 		} else {
// 			r.head = r.head.next
// 		}
// 	}
// 	pre.next = pre.next.next
// }