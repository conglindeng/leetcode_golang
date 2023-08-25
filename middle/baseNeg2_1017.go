package middle

import (
	"bytes"
	"strconv"
)

func BaseNeg2(n int) string {
	nums := make([]int, 0)
	for n != 0 {
		nums = append(nums, n%2)
		n = n / 2
	}
	if len(nums) == 0 {
		nums = append(nums, 0)
	}
	for i := 0; i < len(nums); i++ {
		if i%2 == 1 && nums[len(nums)-1-i] == '1' {
			
		}
	}
	return intArray2StrWithReverse(nums)
}

func convertToBinary(n int) string {
	nums := make([]int, 0)
	for n != 0 {
		nums = append(nums, n%2)
		n = n / 2
	}
	if len(nums) == 0 {
		nums = append(nums, 0)
	}
	return intArray2StrWithReverse(nums)
}

func intArray2StrWithReverse(nums []int) string {
	var b bytes.Buffer
	for i := len(nums) - 1; i >= 0; i-- {
		b.WriteString(strconv.Itoa(nums[i]))
	}
	return b.String()
}
