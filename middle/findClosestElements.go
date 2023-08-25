package middle

import "sort"

func FindClosestElements(arr []int, k int, x int) []int {
	if len(arr) == 1 {
		return arr
	}
	left := getIdx(arr, x)
	right := left + 1
	res := make([]int, 0)
	for k > 0 {
		if left < 0 {
			res = append(res, arr[right])
			right++
		} else if right >= len(arr) {
			res = append([]int{arr[left]}, res...)
			left--
		} else {
			if x-arr[left] > arr[right]-x {
				res = append(res, arr[right])
				right++
			} else {
				res = append([]int{arr[left]}, res...)
				left--
			}
		}
		k--
	}
	return res
}

// first less than or equals num
func getIdx(arr []int, target int) int {
	left, right := 0, len(arr)-1
	var mid int
	for left < right {
		mid = left + (right-left+1)/2
		if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}

func findClosestElements(arr []int, k int, x int) []int {
	sort.SliceStable(arr, func(i, j int) bool {
		return abs(x-arr[i]) < abs(x-arr[j])
	})
	res := arr[:k]
	sort.Ints(res)
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findClosestElements_1(arr []int, k int, x int) []int {
	right := sort.SearchInts(arr, x)
	left := right - 1
	for k > 0 {
		if left < 0 {
			right++
		} else if right >= len(arr) || arr[right]-x >= x-arr[left] {
			left--
		} else {
			right++
		}
		k--
	}
	return arr[left+1 : right]
}
