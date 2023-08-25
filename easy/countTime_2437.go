package easy

import (
	"strconv"
	"strings"
)

func CountTime(time string) int {
	validateTime := func(time string) bool {
		s := strings.Split(time, ":")
		h, _ := strconv.Atoi(s[0])
		m, _ := strconv.Atoi(s[1])
		return h < 24 && m <= 59
	}

	var dfs func(time string, i int)
	res := 0
	dfs = func(time string, i int) {
		if i == 5 {
			if validateTime(time) {
				res++
			}
			return
		}
		if time[i] == '?' {
			for n := 0; n < 10; n++ {
				newTime := time[:i] + strconv.Itoa(n) + time[i+1:]
				dfs(newTime, i+1)
			}
		} else {
			dfs(time, i+1)
		}

	}
	dfs(time, 0)
	return res
}

func countTime(time string) int {
    countHour := 0
    countMinute := 0
    for i := 0; i < 24; i++ {
        hiHour := byte(i / 10)
        loHour := byte(i % 10)
        if (time[0] == '?' || time[0] == hiHour+'0') &&
            (time[1] == '?' || time[1] == loHour+'0') {
            countHour++
        }
    }
    for i := 0; i < 60; i++ {
        hiMinute := byte(i / 10)
        loMinute := byte(i % 10)
        if (time[3] == '?' || time[3] == hiMinute+'0') &&
            (time[4] == '?' || time[4] == loMinute+'0') {
            countMinute++
        }
    }
    return countHour * countMinute
}
