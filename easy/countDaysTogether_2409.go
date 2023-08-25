package easy

import "time"

func CountDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	if leaveBob < arriveAlice || arriveBob > leaveAlice {
		return 0
	}
	maxArrive := arriveAlice
	if arriveBob > maxArrive {
		maxArrive = arriveBob
	}
	minLeave := leaveAlice
	if minLeave > leaveBob {
		minLeave = leaveBob
	}
	layout := "2006-01-02"
	t1, _ := time.Parse(layout, "2023-"+maxArrive)
	t2, _ := time.Parse(layout, "2023-"+minLeave)
	return int(t2.Sub(t1).Hours()/24) + 1
}
