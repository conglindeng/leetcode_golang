package middle

import (
	"fmt"
	"strconv"
)

var count_check = [10]int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}

func rotatedDigits(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		valid, diff := true, false
		for _, c := range strconv.Itoa(i) {
			fmt.Println(c - '0')
			if count_check[c-'0'] == -1 {
				valid = false
			} else if count_check[c-'0'] == 1 {
				diff = true
			}
		}
		if valid && !diff {
			count++
		}
	}
	return count
}
