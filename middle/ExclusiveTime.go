package middle

import (
	"container/list"
	"strconv"
	"strings"
)

const START = "start"

func ExclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	work := list.New()
	for _, item := range logs {
		split := strings.Split(item, ":")
		index, _ := strconv.Atoi(split[0])
		time, _ := strconv.Atoi(split[2])
		if START == split[1] {
			if work.Len() > 0 {
				peek := work.Front().Value.([]int)
				res[peek[0]] += time - peek[1]
				// peek[1]=time
			}
			work.PushFront([]int{index, time})
		} else {
			front := work.Front()
			work.Remove(front)
			last := front.Value.([]int)
			res[last[0]] += time - last[1] + 1
			if work.Len() > 0 {
				peek := work.Front().Value.([]int)
				peek[1] = time + 1
			}
		}
	}
	return res
}