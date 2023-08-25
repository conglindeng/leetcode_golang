package easy

import "sort"

func canBeEqual(target []int, arr []int) bool {
	count1 := map[int]int{}
	for item, _ := range target {
		count1[item] += 1
	}
	count2 := map[int]int{}
	for _, item := range target {
		count2[item] += 1
	}
	for num, count := range count1 {
		if count2[num] != count {
			return false
		}
	}
	return true
}

func canBeEqual_1(target []int, arr []int) bool {
	count := map[int]int{}
	for i := 0; i < len(arr); i++ {
		count[target[i]]++
		count[arr[i]]--
	}
	for _, count := range count {
		if count != 0 {
			return false
		}
	}
	return true
}

func canBeEqual_2(target []int, arr []int) bool {
	sort.Ints(arr)
	sort.Ints(target)
	for i := 0; i < len(arr); i++ {
		if arr[i] != target[i] {
			return false
		}
	}
	return true
}
