package middle

func singleNumber(nums []int) int {
	res := int32(0)
	for i := 0; i < 32; i++ {
		sum := 0
		for _, n := range nums {
			sum += (n >> i) & 1
		}
		if sum%3 > 0 {
			res |= 1 << i
		}
	}
	return int(res)
}
