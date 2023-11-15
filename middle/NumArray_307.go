package middle

type NumArray struct {
	nums, sum []int
}

func Constructor_NumArray(nums []int) NumArray {
	NumArray := NumArray{nums: nums}
	sum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	NumArray.sum = sum
	return NumArray
}

func (this *NumArray) Update(index int, val int) {
	offset := val - this.nums[index]
	this.nums[index] = val
	for i := index + 1; i < len(this.sum); i++ {
		this.sum[i] += offset
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum[right] - this.sum[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
