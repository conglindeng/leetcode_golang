package easy

import (
	"strconv"
	"strings"
)

func AreNumbersAscending(s string) bool {
	split := strings.Split(s, " ")
	pre := -1
	for _, cur := range split {
		if n, err := strconv.Atoi(cur); err == nil && n != 0 {
			if n <= pre {
				return false
			} else {
				pre = n
			}
		}
	}
	return true
}

func AreNumbersAscending_(s string) bool {
	//todo: dcl
	return true
}
