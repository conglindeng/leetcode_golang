package middle

// type node struct{
// 	val int
// 	pre *node
// 	next *node
// }

type MyCircularQueue struct {
	size       int
	capacity   int
	nums       []int
	head, tail int
}

func MyCircularQueue_Constructor(k int) MyCircularQueue {
	// m := new(MyCircularQueue)
	m := MyCircularQueue{capacity: k}
	m.nums = make([]int, k)
	return m
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.size == 0 {
		this.nums[this.tail] = value
	} else {
		this.tail = (this.tail + 1) % this.capacity
		this.nums[this.tail] = value
	}
	this.size += 1
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.tail != this.head {
		this.head = (this.head + 1) % this.capacity
	}
	this.size -= 1
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.nums[this.head]
	}
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.nums[this.tail]
	}
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.capacity == this.size
}
