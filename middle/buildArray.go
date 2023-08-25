package middle

import "reflect"

const (
	STACK_PUSH = "Push"
	STACK_POP  = "Pop"
)

func BuildArray(target []int, n int) []string {
	nums := make([]int, len(target))
	res := make([]string, 0)
	index := 1
	m := 0
	nums[0] = 1
	res = append(res, STACK_PUSH)
	for !reflect.DeepEqual(nums, target) {
		if nums[m] != target[m] {
			index++
			nums[m] = index
			res = append(res, STACK_POP)
			res = append(res, STACK_PUSH)
		} else {
			m++
			index++
			nums[m] = index
			res = append(res, STACK_PUSH)
		}
	}
	return res
}

func buildArray(target []int, n int) []string {
	pre := 0
	res := make([]string, 0)
	for _, n := range target {
		for i := 0; i < n-pre-1; i++ {
			res = append(res, "Push")
			res = append(res, "Pop")
		}
		res = append(res, "Push")
		pre = n
	}
	return res
}
