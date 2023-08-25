package middle

import "sort"

func AdvantageCount(nums1 []int, nums2 []int) []int {
	l := len(nums1)
	idx1 := make([]int, l)
	idx2 := make([]int, l)
	for i := 1; i < l; i++ {
		idx1[i] = i
		idx2[i] = i
	}
	sort.Slice(idx1, func(i, j int) bool { return nums1[idx1[i]] < nums1[idx1[j]] })
	sort.Slice(idx2, func(i, j int) bool { return nums2[idx2[i]] < nums2[idx2[j]] })
	res := make([]int, len(nums1))
	left, right := 0, l-1
	for i := 0; i < l; i++ {
		if nums1[idx1[i]] > nums2[idx2[left]] {
			res[idx2[left]] = nums1[idx1[i]]
			left++
		} else {
			res[idx2[right]] = nums1[idx1[i]]
			right--
		}
	}
	return res
}
