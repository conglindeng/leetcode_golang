package middle

/*

import "github.com/emirpasic/gods/trees/redblacktree"

type StockPrice struct {
	prices   *redblacktree.Tree
	priceMap map[int]int
	maxTime  int
}

func Constructor_StockPrice() StockPrice {
	return StockPrice{maxTime: 0, priceMap: map[int]int{}, prices: redblacktree.NewWithIntComparator()}
}

func (this *StockPrice) Update(timestamp int, price int) {
	this.maxTime = max(timestamp, this.maxTime)
	v, _ := this.priceMap[timestamp]
	if v > 0 {
		if count, _ := this.prices.Get(v); count.(int) > 1 {
			this.prices.Put(v, count.(int)-1)
		} else {
			this.prices.Remove(v)
		}
	}
	this.priceMap[timestamp] = price
	time := 1
	if t, exist := this.prices.Get(price); exist {
		time = t.(int) + 1
	}
	this.prices.Put(price, time)
}

func (this *StockPrice) Current() int {
	return this.priceMap[this.maxTime]
}

func (this *StockPrice) Maximum() int {
	return this.prices.Right().Key.(int)
}

func (this *StockPrice) Minimum() int {
	return this.prices.Left().Key.(int)

}
*/
