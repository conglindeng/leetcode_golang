package middle

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	num := map[int]struct{}{}
	added := map[int]struct{}{}
	res := make([]int, 0)
	for _, v := range nums1 {
		num[v] = struct{}{}
	}
	var curNum map[int]struct{}
	curNum = map[int]struct{}{}
	for _, v := range nums2 {
		if _, duplicate := curNum[v]; !duplicate {
			if _, exsit := num[v]; exsit {
				if _, add := added[v]; !add {
					res = append(res, v)
					added[v] = struct{}{}
				}
			} else {
				num[v] = struct{}{}
			}
			curNum[v] = struct{}{}
		}
	}
	curNum = map[int]struct{}{}
	for _, v := range nums3 {
		if _, duplicate := curNum[v]; !duplicate {
			if _, exsit := num[v]; exsit {
				if _, add := added[v]; !add {
					res = append(res, v)
					added[v] = struct{}{}
				}
			} else {
				num[v] = struct{}{}
			}
			curNum[v] = struct{}{}
		}
	}
	return res
}

func twoOutOfThree_(nums1 []int, nums2 []int, nums3 []int) []int {
	mask := map[int]int{}
	for i, nums := range [][]int{nums1, nums2, nums3} {
		for _, x := range nums {
			mask[x] |= 1 << i
		}
	}
	res := make([]int, 0)
	for k, v := range mask {
		if v&(v-1) > 0 {
			res = append(res, k)
		}
	}
	return res
}
