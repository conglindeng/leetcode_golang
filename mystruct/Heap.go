package mystruct

import (
	"sort"
)

type Heap struct {
	sort.IntSlice
}

func Heap_Init() *Heap {
	return &Heap{sort.IntSlice{}}
}

func (h *Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
	temp := h.IntSlice
	i := temp[len(temp)-1]
	h.IntSlice = temp[:len(temp)-1]
	return i
}
