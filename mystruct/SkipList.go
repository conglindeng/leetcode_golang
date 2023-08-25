package mystruct

import "math/rand"

const (
	pFactor  = 0.25
	maxLevel = 32
)

//why use array

type SkipListNode struct {
	right, down *SkipListNode
	val, level  int
}

type Skiplist struct {
	heads []*SkipListNode
	level int
}

func Skiplist_Init() Skiplist {
	return Skiplist{make([]*SkipListNode, maxLevel), -1}
}

func (this *Skiplist) find(target int) (*SkipListNode, bool) {
	if this.level == -1 {
		return nil, false
	}
	temp := this.heads[this.level-1]
	for temp != nil {
		if temp.right == nil || temp.right.val < target {
			temp = temp.down
		} else if temp.val == target {
			return temp, true
		} else {
			temp = temp.right
		}
	}
	return temp, false
}

func (this *Skiplist) Search(target int) bool {
	_, ok := this.find(target)
	return ok
}

func (this *Skiplist) Add(num int) {
	node, ok := this.find(num)
	if !ok {
		//exist,do nothing
		return
	}
	//insert it on level 0
	newNode := &SkipListNode{val: num, level: 0, right: node.right}
	node.right = newNode
	for rand.Float32() < pFactor {
		
	}
}

//same as delete
func (this *Skiplist) Erase(num int) bool {
	node, ok := this.find(num)
	if !ok {
		return false
	}
	level := node.level
	for i := level; i < this.level; i++ {
		h := this.heads[i]
		for h.right != nil {
			if h.right.val == num {
				h.right = h.right.right
				break
			}
		}
		h = this.heads[i]
		if h.right == nil {
			this.level--
		}
	}
	return true
}
