package middle

func MinElements(nums []int, limit int, goal int) int {
	sum := 0
	for _,i := range nums {
		sum += i
	}
	diff := abs(goal - sum)
	return (diff+limit-1)/limit
}
