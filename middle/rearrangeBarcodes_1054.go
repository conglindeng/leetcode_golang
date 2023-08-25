package middle

import "sort"

func RearrangeBarcodes(barcodes []int) []int {
	count := map[int]int{}
	res := make([]int, len(barcodes))
	for _, n := range barcodes {
		count[n]++
	}

	keys := make([]int, 0, len(count))

	for key := range count {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return count[keys[i]] > count[keys[j]]
	})
	need, index, needIndex := 0, 0, 1

	for _, k := range keys {
		v := count[k]
		if v <= need {
			for i := 0; i < v; i++ {
				res[needIndex] = k
				needIndex += 2
			}
			need = need - v
			//cur nums can fill positions for sperating every num
		} else {
			for i := 0; i < need; i++ {
				res[needIndex] = k
				needIndex += 2
			}
			needIndex = index + 1
			for i := 0; i < v-need; i++ {
				res[index] = k
				if i == v-need-1 {
					index += 1
				} else {
					index += 2
				}
			}
			need = v - need - 1
		}
	}
	return res
}
