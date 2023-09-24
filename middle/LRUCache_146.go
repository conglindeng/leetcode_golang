package middle

type LRUCacheNode struct {
	pre, next *LRUCacheNode
	key, val  int
}

type LRUCache struct {
	size, capacity int
	cache          map[int]*LRUCacheNode
	head, tail     *LRUCacheNode
}

func LRUCache_Constructor(capacity int) LRUCache {
	head, tail := &LRUCacheNode{}, &LRUCacheNode{}
	head.next = tail
	tail.pre = head
	return LRUCache{size: 0, capacity: capacity, head: head, tail: tail, cache: map[int]*LRUCacheNode{}}
}

func (this *LRUCache) Get(key int) int {
	if v, exist := this.cache[key]; exist {

		//disconnect pointer relationship
		pre := v.pre
		pre.next = v.next
		v.next.pre = pre

		//move to head
		this.head.next.pre = v
		v.next = this.head.next
		this.head.next = v
		v.pre = this.head

		return v.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, exist := this.cache[key]; exist {

		//disconnect pointer relationship
		pre := v.pre
		pre.next = v.next
		v.next.pre = pre

		//move to head
		this.head.next.pre = v
		v.next = this.head.next
		this.head.next = v
		v.pre = this.head

		v.val = value
		return
	}
	if this.size == this.capacity {
		//remove lasted used k-v pair
		pre := this.tail.pre
		pre.pre.next = this.tail
		this.tail.pre = pre.pre
		this.size--
		delete(this.cache, pre.key)
	}
	newNode := &LRUCacheNode{key: key, val: value}
	this.cache[key] = newNode
	this.head.next.pre = newNode
	newNode.pre = this.head
	newNode.next = this.head.next
	this.head.next = newNode
	this.size++
}
