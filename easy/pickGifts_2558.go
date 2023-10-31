package easy

import (
	"container/heap"
	"math"
)

func PickGifts(gifts []int, k int) int64 {
	res := int64(0)
	work := &giftHP{}
	for _, v := range gifts {
		heap.Push(work, v)
	}
	for i := 0; i < k; i++ {
		cur := heap.Pop(work).(int)
		if cur == 1 {
			heap.Push(work, 1)
			break
		}
		a := int(math.Sqrt(float64(cur)))
		heap.Push(work, a)
	}
	for _, v := range *work {
		res += int64(v)
	}
	return res
}

type giftHP []int

func (h giftHP) Len() int           { return len(h) }
func (h giftHP) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h giftHP) Less(i, j int) bool { return h[i] > h[j] }

func (h *giftHP) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *giftHP) Pop() any {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
