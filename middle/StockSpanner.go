package middle

import "math"

type StockSpanner struct {
	stack [][2]int
	idx   int
}

func StockSpanner_Constructor() StockSpanner {
	return StockSpanner{[][2]int{{-1, math.MaxInt}}, -1}
}

func (this *StockSpanner) Next(price int) int {
	this.idx++
	for this.stack[len(this.stack)-1][1] <= price {
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, [2]int{this.idx, price})
	return this.idx - this.stack[len(this.stack)-2][0]
}
