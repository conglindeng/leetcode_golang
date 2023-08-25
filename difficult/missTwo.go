package difficult

func missingTwo(nums []int) []int {
	m := 0
	n := len(nums) + 2
	for _, num := range nums {
		m ^= num
	}
	for i := 1; i <= n; i++ {
		m ^= i
	}
	lsb := m & -m
	type1, type2 := 0, 0
	for _, num := range nums {
		if num&lsb > 1 {
			type1 ^= num
		} else {
			type2 ^= num
		}
	}
	for i := 1; i <= n; i++ {
		if i&lsb > 1 {
			type1 ^= i
		} else {
			type2 ^= i
		}
	}
	return []int{type1, type2}
}
