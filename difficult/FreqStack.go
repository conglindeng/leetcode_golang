package difficult

type FreqStack struct {
	maxFreq  int
	stackArr [][]int
	freq     map[int]int
}

func FreqStack_Constructor() FreqStack {
	return FreqStack{0, [][]int{}, map[int]int{}}
}

func (this *FreqStack) Push(val int) {
	this.freq[val]++
	this.maxFreq = max(this.maxFreq, this.freq[val])
	for len(this.stackArr) <= this.maxFreq {
		this.stackArr = append(this.stackArr, []int{})
	}
	this.stackArr[this.freq[val]] = append(this.stackArr[this.freq[val]], val)
}

func (this *FreqStack) Pop() int {
	v := this.stackArr[this.maxFreq][len(this.stackArr[this.maxFreq])-1]
	this.stackArr[this.maxFreq] = this.stackArr[this.maxFreq][:len(this.stackArr[this.maxFreq])-1]
	if len(this.stackArr[this.maxFreq]) == 0 {
		this.maxFreq--
	}
	this.freq[v]--
	return v
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
