package middle

import "container/heap"

type HalveArray []float64

func (h HalveArray) Len() int {
	return len(h)
}

func (h HalveArray) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h HalveArray) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HalveArray) Push(x any) {
	*h = append(*h, x.(float64))
}
func (h *HalveArray) Pop() any {
	m := *h
	x := m[len(m)-1]
	*h = m[:len(m)-1]
	return x
}

func halveArray(nums []int) int {
	h := HalveArray{}
	sum := 0.0
	for _, v := range nums {
		x := float64(v)
		sum += x
		h = append(h, x)
	}
	heap.Init(&h)
	res := 0
	sub := 0.0
	for {
		if sub >= sum/2 {
			break
		}
		res++
		a := heap.Pop(&h).(float64)
		sub += a / 2
		heap.Push(&h, a/2)
	}
	return res
}
