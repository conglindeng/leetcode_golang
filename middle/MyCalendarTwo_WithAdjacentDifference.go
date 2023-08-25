package middle

import "github.com/emirpasic/gods/trees/redblacktree"

type MyCalendarTwo struct {
	*redblacktree.Tree
}

func Constructor1111() MyCalendarTwo {
	return MyCalendarTwo{redblacktree.NewWithIntComparator()}
}

func (it *MyCalendarTwo) Book(start int, end int) bool {
	it.add(start, 1)
	it.add(end, -1)
	ite := it.Iterator()
	count := 0
	for ite.Next() {
		count += ite.Value().(int)
		if count > 2 {
			it.add(start, -1)
			it.add(end, 1)
			return false
		}
	}
	return true
}

func (it *MyCalendarTwo) add(key int, val int) {
	v, ok := it.Get(key)
	if ok {
		it.Put(key, val+v.(int))
	} else {
		it.Put(key, val)
	}
}