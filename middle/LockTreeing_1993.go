package middle

type InnerLockingTree struct {
	childs          []*InnerLockingTree
	parent          *InnerLockingTree
	val, lockedUser int // -1 - not locked
}

type LockingTree struct {
	LockingTreeMap map[int]*InnerLockingTree
}

func LockingTree_Constructor(parent []int) LockingTree {
	tmp := map[int]*InnerLockingTree{}
	childsMap := map[int][]int{}
	ma := -1
	for i, k := range parent {
		childsMap[k] = append(childsMap[k], i)
		ma = max(ma, k)
	}
	var buildLockingTree func(val int) *InnerLockingTree
	buildLockingTree = func(val int) *InnerLockingTree {
		cur := &InnerLockingTree{lockedUser: -1, val: val}
		tmp[val] = cur
		for _, v := range childsMap[val] {
			x := buildLockingTree(v)
			cur.childs = append(cur.childs, x)
			x.parent = cur
		}
		return cur
	}
	buildLockingTree(0)
	return LockingTree{LockingTreeMap: tmp}
}

func (this *LockingTree) Lock(num int, user int) bool {
	ilt := this.LockingTreeMap[num]
	if ilt.lockedUser == -1 {
		ilt.lockedUser = user
		return true
	}
	return false
}

func (this *LockingTree) Unlock(num int, user int) bool {
	ilt := this.LockingTreeMap[num]
	if ilt.lockedUser == user {
		ilt.lockedUser = -1
		return true
	}
	return false
}

func (this *LockingTree) Upgrade(num int, user int) bool {
	ilt := this.LockingTreeMap[num]
	if ilt.lockedUser != -1 {
		return false
	}
	tmp := make([]*InnerLockingTree, 0)
	tmp = append(tmp, ilt.childs...)
	lockedChilds := []*InnerLockingTree{}
	for len(tmp) > 0 {
		first := tmp[0]
		tmp = tmp[1:]
		if first.lockedUser != -1 {
			lockedChilds = append(lockedChilds, first)
		}
		tmp = append(tmp, first.childs...)
	}
	if len(lockedChilds) == 0 {
		return false
	}
	pa := ilt.parent
	for pa != nil {
		if pa.lockedUser != -1 {
			return false
		}
		pa = pa.parent
	}
	for _, ilt2 := range lockedChilds {
		ilt2.lockedUser = -1
	}
	ilt.lockedUser = user
	return true
}

/**
 * Your LockingTree object will be instantiated and called as such:
 * obj := Constructor(parent);
 * param_1 := obj.Lock(num,user);
 * param_2 := obj.Unlock(num,user);
 * param_3 := obj.Upgrade(num,user);
 */
