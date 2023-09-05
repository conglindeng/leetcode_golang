package easy

import "sort"

func minNumber(nums1 []int, nums2 []int) int {
	exsit := map[int]bool{}
	m1 := nums1[0]
	m2 := nums2[0]
	for _, v := range nums1 {
		exsit[v] = true
		m1 = min(m1, v)
	}
	sort.Ints(nums2)
	for _, v := range nums2 {
		if exsit[v] {
			return v
		}
		m2 = min(m2, v)
	}
	return min(m1*10+m2, m2*10+m1)
}
