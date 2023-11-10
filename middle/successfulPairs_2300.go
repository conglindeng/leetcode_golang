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
	sort.Ints(potions)
	tmp := [][2]int{}
	for i, v := range spells {
		tmp = append(tmp, [2]int{i, v})
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i][1] > tmp[j][1]
	})
	spellIdx, potionIdx := 0, 0
	res := make([]int, len(spells))
	for spellIdx < len(spells) {
		for potionIdx < len(potions) && tmp[spellIdx][1]*potions[potionIdx] < int(success) {
			potionIdx++
		}
		res[tmp[spellIdx][0]] = len(potions) - potionIdx
		spellIdx++
	}
	return res
}
