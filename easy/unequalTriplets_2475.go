package easy

import "sort"

func unequalTriplets(nums []int) int {
	count := 0
	n := len(nums)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for m := j + 1; j < n; j++ {
				if nums[m] != nums[j] && nums[j] != nums[i] {
					count++
				}
			}
		}
	}
	return count
}

func unequalTriplets_sort(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	res := 0
	for i, j := 0, 0; i < n; i = j {
		for j < n && nums[j] == nums[i] {
			j++
		}
		res += i * (j - i) * (n - j)
	}
	return res
}

func unequalTriplets_hash(nums []int) int {
	count := map[int]int{}
	for _, num := range nums {
		count[num]++
	}
	res, pre := 0, 0
	n := len(nums)
	for _, c := range count {
		res += pre * c * (n - pre - c)
		pre = pre + c
	}
	return res
}
