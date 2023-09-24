package middle

func minCapability(nums []int, k int) int {
	mi, ma := nums[0], nums[0]
	for _, x := range nums {
		if mi > x {
			mi = x
		}
		if ma < x {
			ma = x
		}
	}
	for mi <= ma {
		mid := mi + (ma-mi)/2
		count := 0
		visited := false
		for _, x := range nums {
			if x <= mid && !visited {
				count++
				visited = true
			} else {
				visited = false
			}
		}
		if count >= k {
			ma = mid - 1
		} else {
			mi = mid + 1
		}
	}
	return mi
}
