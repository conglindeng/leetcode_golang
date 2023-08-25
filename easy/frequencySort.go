package easy

import (
	"sort"
)

func FrequencySort(nums []int) []int {
	count := map[int]int{}
	for _, num := range nums {
		if _, ok := count[num]; !ok {
			count[num] = 1
		} else {
			count[num]++
		}
	}
	temp := make([]int, 0, len(count))
	for key, _ := range count {
		temp = append(temp, key)
	}
	sort.SliceStable(temp, func(i, j int) bool {
		if count[temp[i]] != count[temp[j]] {
			return count[temp[i]] > count[temp[j]]
		} else {
			return temp[i] < temp[j]
		}
	})
	res := make([]int, 0, len(nums))
	for i := len(temp) - 1; i >= 0; i-- {
		for j := 0; j < count[temp[i]]; j++ {
			res = append(res, temp[i])
		}
	}
	return res
}

func FrequencySort_(nums []int) []int {
	count := map[int]int{}
	for _, i := range nums {
		count[i]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if count[nums[i]] != count[nums[j]] {
			return count[nums[i]] > count[nums[j]]
		} else {
			return nums[i] > nums[j]
		}
	})
	return nums
}
