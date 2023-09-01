package middle

func rotate(nums []int, k int) {
	doRotate := func(nums []int, left, right int) {
		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	k = k%len(nums) - 1
	doRotate(nums, 0, len(nums)-1)
	doRotate(nums, 0, k)
	doRotate(nums, k+1, len(nums)-1)
}

func rotate_2(nums []int, k int) {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[(i+k)%len(nums)] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = res[i]
	}
}

func rotate_3(nums []int, k int) {

}
