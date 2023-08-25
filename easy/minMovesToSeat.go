package easy

import "sort"

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	res := 0
	for i := 0; i < len(seats); i++ {
		res += abs(seats[i] - students[i])
	}
	return res
}
