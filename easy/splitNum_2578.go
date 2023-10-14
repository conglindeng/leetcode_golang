package easy

import "sort"

func SplitNum(num int) int {
	process := make([]int, 0)
	for num > 0 {
		process = append(process, num%10)
		num = num / 10
	}
	if len(process)%2 != 0 {
		process = append(process, 0)
	}
	sort.Ints(process)
	carry := 0
	first, second := len(process)-2, len(process)-1
	res := 0
	shift := 1
	for first >= 0 {
		cur := process[first] + process[second] + carry
		res = cur%10*shift + res
		carry = cur / 10
		shift *= 10
		first -= 2
		second -= 2
	}
	if carry != 0 {
		res = carry*shift + res
	}
	return res
}
