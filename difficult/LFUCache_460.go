package difficult

type LFUCacheNode struct {
	pre, next, up, down *LFUCacheNode
	key, val, frequence int
}

type LFUCache struct {
	size, capacity int
	head, tail     *LFUCacheNode
	cache          map[int]*LFUCacheNode
}

func Constructor(capacity int) LFUCache {
	head, tail := &LFUCacheNode{}, &LFUCacheNode{}
	head.next = tail
	tail.pre = head
	return LFUCache{size: 0, capacity: capacity, head: head, tail: tail, cache: map[int]*LFUCacheNode{}}
}

func (this *LFUCache) Get(key int) int {
	if v, exist := this.cache[key]; exist {
		v.frequence++
		//find up
		up := v.findUp()
		pre := up.pre
		if up.key == v.key {
			//remove cur node
			p := up.pre
			p.next = up.next
			up.next.pre = p
		} else {
			u := v.up
			u.down = v.down
			if v.down != nil {
				v.down.up = u
			}
		}
		v.pre = nil
		v.next = nil
		v.up = nil
		v.down = nil
		if pre.frequence == v.frequence {
			v.next = pre.next
			v.pre = pre.pre
			pre.pre.next = v
			pre.next.pre = v

			pre.next = nil
			pre.pre = nil

			v.down = pre
			pre.up = v

		} else {
			v.next = pre.next
			pre.next.pre = v
			v.pre = pre
			pre.next = v
		}
		return v.val
	}
	return -1
}

func (this *LFUCacheNode) findUp() *LFUCacheNode {
	cur := this
	for cur.up != nil {
		cur = cur.up
	}
	return cur
}

func (this *LFUCache) Put(key int, value int) {
	if v, exist := this.cache[key]; exist {
		v.frequence++
		v.val = value
		//find up
		up := v.findUp()
		pre := up.pre
		if up.key == v.key {
			if up.down == nil {
				// remove cur node
				p := up.pre
				p.next = up.next
				up.next.pre = p
			} else {
				pre.next = up.down
				up.down.pre = pre
				up.down.next = up.next
				up.next.pre = up.down
				up.down.up = nil
			}
		} else {
			u := v.up
			u.down = v.down
			if v.down != nil {
				v.down.up = u
			}
		}
		v.pre = nil
		v.next = nil
		v.up = nil
		v.down = nil
		if pre.frequence == v.frequence {
			v.next = pre.next
			v.pre = pre.pre
			pre.pre.next = v
			pre.next.pre = v

			pre.next = nil
			pre.pre = nil

			v.down = pre
			pre.up = v

		} else {
			v.next = pre.next
			pre.next.pre = v
			v.pre = pre
			pre.next = v
		}
	} else {
		if this.size == this.capacity {
			ln := this.tail.pre
			for ln.down != nil {
				ln = ln.down
			}
			if ln == this.tail.pre {
				pre := ln.pre
				pre.next = this.tail
				this.tail.pre = pre

				ln.pre = nil
				ln.next = nil

			} else {
				up := ln.up
				up.down = nil
				ln.up = nil
			}
			this.size--
			delete(this.cache, ln.key)
		}
		v := &LFUCacheNode{frequence: 1, key: key, val: value}
		pre := this.tail.pre
		if pre.frequence == v.frequence {
			v.next = pre.next
			v.pre = pre.pre
			pre.pre.next = v
			pre.next.pre = v

			pre.next = nil
			pre.pre = nil

			v.down = pre
			pre.up = v

		} else {
			v.next = pre.next
			pre.next.pre = v
			v.pre = pre
			pre.next = v
		}
		this.size++
		this.cache[key] = v
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
