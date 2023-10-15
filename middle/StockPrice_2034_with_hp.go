package middle

import "container/heap"

type StockPrice struct {
	m, s     hp
	priceMap map[int]int
	maxTime  int
}

func Constructor_StockPrice() StockPrice {
	return StockPrice{priceMap: map[int]int{}, maxTime: 0, m: hp{}, s: hp{}}
}

func (this *StockPrice) Update(timestamp int, price int) {
	this.priceMap[timestamp] = price
	if timestamp > this.maxTime {
		this.maxTime = timestamp
	}
	heap.Push(&this.m, pair{price: price, time: timestamp})
	heap.Push(&this.s, pair{price: -price, time: timestamp})
}

func (this *StockPrice) Current() int {
	return this.priceMap[this.maxTime]
}

func (this *StockPrice) Maximum() int {
	for {
		if first:=this.m.
	}
}

func (this *StockPrice) Minimum() int {

}

type pair struct{ time, price int }

type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h hp) Less(i, j int) bool { return h[i].price < h[j].price }

func (h *hp) Push(x any) {
	*h = append(*h, x.(pair))
}
func (h *hp) Pop() any {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
