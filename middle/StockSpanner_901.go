package middle

import "math"

type StockSpanner struct {
	stack [][2]int
	size  int
}

func StockSpanner_Constructor() StockSpanner {
	return StockSpanner{stack: [][2]int{{0, math.MinInt}}}
}

func (this *StockSpanner) Next(price int) int {
	this.size++
	for len(this.stack) > 1 && this.stack[len(this.stack)-1][1] <= price {
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, [2]int{this.size, price})
	return this.size - this.stack[len(this.stack)-2][0]
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
