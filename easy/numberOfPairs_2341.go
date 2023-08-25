package easy


func NumberOfPairs(nums []int) []int {
	count := map[int]int{}
	for _, v := range nums {
		count[v]++
	}
	pair, left := 0, 0
	for _, v := range count {
		pair += v / 2
		left += v % 2
	}
	return []int{pair, left}
}

func NumberOfPairs_(nums []int) []int {
	pair := 0
	exsit := map[int]bool{}
	for _, num := range nums {
		exsit[num] = !exsit[num]
		if !exsit[num] {
			pair++
		}
	}
	return []int{pair, len(nums) - 2*pair}
}
