package middle

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	res := make([]int, len(spells))
	for i, spell := range spells {
		left, right := 0, len(potions)
		for left < right {
			mid := (right-left)>>1 + left
			if potions[mid]*spell < int(success) {
				left = mid + 1
			} else {
				right = mid
			}
		}
		res[i] = len(potions) - left
	}
	return res
}

func successfulPairs_double_pointer(spells []int, potions []int, success int64) []int {

}
