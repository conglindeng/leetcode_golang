package middle

type node struct {
	val  int
	pre  *node
	next *node
}

type MyCircularDeque struct {
	capacity   int
	size       int
	head, tail *node
}

func Constructor_(k int) MyCircularDeque {
	return MyCircularDeque{capacity: k}
}

func (d *MyCircularDeque) InsertFront(value int) bool {
	if d.IsFull() {
		return false
	}
	n := &node{val: value}
	if d.size == 0 {
		d.head = n
		d.tail = n
		d.head.next = d.tail
		d.tail.pre = d.head
	} else {
		n.next = d.head
		d.head.pre = n
		d.head = n
	}
	d.size++
	return true
}

func (d *MyCircularDeque) InsertLast(value int) bool {
	if d.IsFull() {
		return false
	}
	n := &node{val: value}
	if d.size == 0 {
		d.head = n
		d.tail = n
		d.head.next = d.tail
		d.tail.pre = d.head
	} else {
		n.pre = d.tail
		d.tail.next = n
		d.tail = n
	}
	d.size++
	return true

}

func (d *MyCircularDeque) DeleteFront() bool {
	if d.IsEmpty() {
		return false
	}
	if d.size == 1 {
		d.head = nil
		d.tail = nil
	} else {
		head := d.head
		head.next.pre = nil
		d.head = head.next
		head.next = nil

	}
	d.size--
	return true
}

func (d *MyCircularDeque) DeleteLast() bool {
	if d.IsEmpty() {
		return false
	}
	if d.size == 1 {
		d.head = nil
		d.tail = nil
	} else {
		tail := d.tail
		tail.pre.next = nil
		d.tail = tail.pre
		tail.pre = nil
	}
	d.size--
	return true
}

func (d *MyCircularDeque) GetFront() int {
	if d.IsEmpty() {
		return -1
	}
	return d.head.val
}

func (d *MyCircularDeque) GetRear() int {
	if d.IsEmpty() {
		return -1
	}
	return d.tail.val
}

func (d *MyCircularDeque) IsEmpty() bool {
	return d.size == 0
}

func (d *MyCircularDeque) IsFull() bool {
	return d.size == d.capacity
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
