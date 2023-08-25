package middle

type linkedListNode struct {
	val       int
	next, pre *linkedListNode
}

type MyLinkedList struct {
	size       int
	tail, head *linkedListNode
}

func MyLinkedList_Constructor() MyLinkedList {
	return MyLinkedList{size: 0}
}

func (this *MyLinkedList) Get(index int) int {
	if this.size == 0 || index >= this.size {
		return -1
	}
	return this.getByIndex(index).val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := linkedListNode{val: val}
	if this.size == 0 {
		this.head = &node
		this.tail = &node
		// this.head.next = this.tail
		// this.tail.pre = this.head
	} else {
		node.next = this.head
		this.head.pre = &node
		this.head = &node
	}
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := linkedListNode{val: val}
	if this.size == 0 {
		this.head = &node
		this.tail = &node
		// this.head.next = this.tail
		// this.tail.pre = this.head
	} else {
		node.pre = this.tail
		this.tail.next = &node
		this.tail = &node
	}
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index <= 0 {
		this.AddAtHead(val)
	} else if index == this.size {
		this.AddAtTail(val)
	} else {
		node := this.getByIndex(index - 1)
		newNode := linkedListNode{val: val, pre: node, next: node.next}
		node.next.pre = &newNode
		node.next = &newNode
		this.size++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	if this.size == 1 {
		this.head = nil
		this.tail = nil
	} else {
		node := this.getByIndex(index)
		if node == this.head {
			next := this.head.next
			this.head.next = nil
			next.pre = nil
			this.head = next
		} else if node == this.tail {
			pre := this.tail.pre
			pre.next = nil
			this.tail.pre = nil
			this.tail = pre
		} else {
			node.next.pre = node.pre
			node.pre.next = node.next
		}
	}
	this.size--
}

func (this *MyLinkedList) getByIndex(index int) *linkedListNode {
	mid := (this.size - 1) / 2
	if mid < index {
		//the target is in the second half, so traveling from the back to the front
		end := this.size - 1
		temp := this.tail
		for index < end {
			temp = temp.pre
			end--
		}
		return temp
	} else {
		// traveling from the front to the back\
		start := 0
		temp := this.head
		for start < index {
			temp = temp.next
			start++
		}
		return temp
	}
}
